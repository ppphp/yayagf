package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	os.Args = []string{"yayagf", "version"}
	t.Main()
}
