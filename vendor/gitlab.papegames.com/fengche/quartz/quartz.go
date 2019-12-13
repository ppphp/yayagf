package quartz

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Quartz struct {
	ctx      context.Context
	cancel   context.CancelFunc
	root     string
	lock     sync.Mutex
	interval time.Duration
	Event    chan string
}

func (q *Quartz) Begin() error {
	if q.cancel != nil {
		return errors.New("quartz has begun")
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	ctx, cancel := context.WithCancel(q.ctx)
	q.cancel = cancel
	q.Event = make(chan string)
	go q.loop(ctx)
	return nil
}

func (q *Quartz) loop(ctx context.Context) {
	t := time.NewTimer(q.interval)
	lastUpdate := time.Now()
	updated := true
	for {
		select {
		case <-t.C:
			filepath.Walk(q.root, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}
				if info.Name() == ".git" {
					return nil
				}
				if info.ModTime().After(lastUpdate) {
					updated = true
					lastUpdate = info.ModTime()
				}
				return nil
			})
			if updated {
				q.Event <- "some update"
			}
			updated = false
			t.Reset(q.interval)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}
}

func (q *Quartz) Stop() {
	if q.cancel != nil {
		q.cancel()
		q.cancel = nil
	}
}

func NewQuartz(dir string, interval time.Duration) (*Quartz, error) {
	mydir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	if s, err := os.Stat(mydir); err != nil {
		return nil, err
	} else if !s.IsDir() {
		return nil, errors.New("should be a directory")
	}

	q := &Quartz{root: mydir, interval: interval, ctx: context.Background()}
	return q, nil
}
