// A file watcher
package watcher

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var ignoreDirs = []string{".git", "node_modules", "vendor"}

func isIgnoreFile(value string) bool {
	if !strings.HasSuffix(value, ".go") {
		return true
	} else if strings.HasPrefix(value, ".") || strings.HasPrefix(value, "#") || strings.HasPrefix(value, "~") {
		return true
	}
	return false
}

func isIgnoreDir(value string) bool {
	for _, i := range ignoreDirs {
		if value == i {
			return true
		}
	}
	return false
}

type Event struct {
	Action   Action
	FilePath string
}

type Action int8

const (
	Create Action = iota
	Update
	Delete
)

type Watcher struct {
	ctx      context.Context
	cancel   context.CancelFunc
	root     string
	lock     sync.Mutex
	interval time.Duration
	Event    chan Event
	Once     sync.Once
}

func (q *Watcher) Begin() error {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Once.Do(func() {
		ctx, cancel := context.WithCancel(q.ctx)
		q.cancel = cancel
		q.Event = make(chan Event)
		go q.loop(ctx)
	})
	return nil
}

func (q *Watcher) loop(ctx context.Context) {
	t := time.NewTimer(q.interval)
	lastUpdate := time.Now()
	updated := true
	var lastSize int64
	for {
		select {
		case tk := <-t.C:
			var size int64
			updatedPath := ""
			if err := filepath.Walk(q.root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}
				if updated {
					return nil
				}
				if info.IsDir() {
					if isIgnoreDir(info.Name()) {
						return filepath.SkipDir
					}
					return nil
				}
				if isIgnoreFile(info.Name()) {
					return nil
				}
				size += info.Size()

				if info.ModTime().After(lastUpdate) {
					updated = true
					updatedPath = path
					lastUpdate = tk
				}
				return nil
			}); err != nil {
				log.Printf("%v\n", err.Error())
			}
			if updated {
				q.Event <- Event{Action: Update, FilePath: updatedPath}
			} else if size < lastSize {
				q.Event <- Event{Action: Delete, FilePath: ""}
			}
			updated = false
			lastSize = size
			t.Reset(q.interval)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}
}

func (w *Watcher) Stop() {
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
		w.Once = sync.Once{}
	}
}

func NewWatcher(dir string, interval time.Duration) (*Watcher, error) {
	mydir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	if s, err := os.Stat(mydir); err != nil {
		return nil, err
	} else if !s.IsDir() {
		return nil, errors.New("should be a directory")
	}

	q := &Watcher{root: mydir, interval: interval, ctx: context.Background()}
	return q, nil
}
