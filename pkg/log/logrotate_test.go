package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewRotationWriter(t *testing.T) {

	_, err := NewRotationWriter("/", HundredMilliSecond)
	require.Error(t, err)

	r, err := NewRotationWriter("./testdata/log", HundredMilliSecond)
	require.NoError(t, err)
	_, err = r.Write([]byte("aaa\n"))
	require.NoError(t, err)
	time.Sleep(time.Second)
}
