package middleware

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntries(t *testing.T) {
	require.NotNil(t, entries())
}

func TestDebugf(t *testing.T) {
	Debugf("%v", "a")
}

func TestPrint(t *testing.T) {
	Print("%v", "a")
}

func TestWarnf(t *testing.T) {
	Warnf("%v", "a")
}

func TestErrorf(t *testing.T) {
	Errorf("%v", "a")
}

func TestInfof(t *testing.T) {
	Infof("%v", "a")
}

func TestGetLogger(t *testing.T) {
	require.NotNil(t, GetLogger())
}

func TestTweak(t *testing.T) {
	Tweak(Config{})
}

func TestGinrus(t *testing.T) {
	Ginrus(GetLogger())
}
