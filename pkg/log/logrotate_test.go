package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewRotationWriter(t *testing.T) {

	_, err := NewRotationWriter("/", NextNsecond)
	require.Error(t, err)

	r, err := NewRotationWriter("./log", NextNsecond)
	require.NoError(t, err)
	_, err = r.Write([]byte("aaa\n"))
	require.NoError(t, err)
}

func TestNextHour(t *testing.T) {
	require.NotNil(t, NextHour(time.Now()))
}
