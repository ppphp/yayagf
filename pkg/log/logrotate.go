package log

import (
	"fmt"
	"os"
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

func (r *RotationWriter) Rotate(now time.Time, next func(time.Time) <-chan time.Time) {
	timer := next(now)
	for t := range timer {
		_ = os.Rename(r.path, fmt.Sprintf("%s_%04d%02d%02d-%02d%02d%02d", r.path, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
		f, err := os.OpenFile(r.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err == nil {
			fo := r.file
			r.rwm.Lock()
			r.file = f
			r.rwm.Unlock()
			_ = fo.Close()
		}
		timer = next(t)
	}
}

// generate a rotation writer, name is the prefix
func NewRotationWriter(name string, next func(time.Time) <-chan time.Time) (*RotationWriter, error) {
	r := &RotationWriter{path: name}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	r.file = f

	go r.Rotate(time.Now(), next)

	return r, nil
}

func NextHour(now time.Time) <-chan time.Time {
	next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())
	return time.After(next.Sub(now))
}

// mostly for testing
func NextNsecond(now time.Time) <-chan time.Time {
	next := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 1, now.Location())
	return time.After(next.Sub(now))
}
