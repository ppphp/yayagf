package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// a rotatable io.Writer
type RotationWriter struct {
	file *os.File
	path string
	rwm  sync.RWMutex
}

func (r *RotationWriter) Write(p []byte) (n int, err error) {
	r.rwm.Lock()
	defer r.rwm.Unlock()
	return r.file.Write(p)
}

func (r *RotationWriter) Rotate(timer Timer) {
	timer1 := time.NewTimer(time.Until(timer.first()))
	for t := range timer1.C {
		// create new
		_ = os.Rename(r.path, fmt.Sprintf("%s_%04d%02d%02d%02d%02d%02d", r.path, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
		f, err := os.OpenFile(r.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err == nil {
			fo := r.file
			r.rwm.Lock()
			r.file = f
			r.rwm.Unlock()
			_ = fo.Close()
		}
		timer1.Reset(timer.tick)

		// delete old, less than deletion prefix
		dir, file := filepath.Split(r.path)
		d := t.Add(-timer.delete)
		deletionPrefix := fmt.Sprintf("%s_%04d%02d%02d%02d%02d%02d", file, d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second())
		files, err := ioutil.ReadDir(dir)
		if err == nil {
			for _, f := range files {
				if f.Name() != file && strings.HasPrefix(f.Name(), file) && f.Name() < deletionPrefix {
					_ = os.Remove(filepath.Join(dir, f.Name()))
				}
			}
		}
	}
}

// generate a rotation writer, name is the prefix
func NewRotationWriter(name string, timer Timer) (*RotationWriter, error) {
	r := &RotationWriter{path: name}

	dir := filepath.Dir(name)
	_ = os.MkdirAll(dir, 0755)
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	r.file = f

	go r.Rotate(timer)

	return r, nil
}

type Timer struct {
	first  func() time.Time
	tick   time.Duration
	delete time.Duration
}

var Hour = Timer{
	first: func() time.Time {
		now := time.Now()
		return time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())
	},
	tick:   time.Hour,
	delete: 3 * 24 * time.Hour,
}

var HundredMilliSecond = Timer{
	first: func() time.Time {
		now := time.Now()
		return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())
	},
	tick:   100 * time.Millisecond,
	delete: 2 * time.Second,
}
