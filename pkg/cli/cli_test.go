package cli

import "testing"

func TestApp(t *testing.T) {
	a := &App{"test", &Command{}}
	a.PrintMeta()
}
