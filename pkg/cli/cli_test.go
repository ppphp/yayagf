package cli

import "testing"

func TestNewApp(t *testing.T) {
	NewApp("test", "")
}

func TestApp(t *testing.T) {
	a := NewApp("test", "")
	a.PrintMeta()
}
