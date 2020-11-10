package watcher

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWatcher(t *testing.T) {
	_, err := NewWatcher("./testdata/.gitignore", time.Millisecond)
	require.Error(t, err)
	_, err = NewWatcher("./testdata/aa", time.Millisecond)
	require.Error(t, err)

	q, err := NewWatcher("./testdata/", time.Millisecond)
	require.NoErrorf(t, err, "cannot create directory: %v", err)

	err = q.Begin()
	require.NoErrorf(t, err, "begin error: %v", err)
	defer q.Stop()

	go func() {
		for s := range q.Event {
			t.Logf("%v", s)
		}
	}()
	for i := 0; i < 5; i++ {
		if f, err := os.OpenFile("./testdata/a", os.O_CREATE|os.O_WRONLY, 0777); err != nil {
			require.NoErrorf(t, err, "open a error: %v", err)
		} else {
			time.Sleep(2 * time.Microsecond)
			_, err := f.Write([]byte("aa"))
			require.NoErrorf(t, err, "write a error: %v", err)
			f.Close()
			time.Sleep(2 * time.Microsecond)
		}
	}
	if err := os.Remove("./testdata/a"); err != nil {
		require.NoErrorf(t, err, "cannot remove: %v", err)
	}
}
