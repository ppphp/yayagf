package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Args = []string{"yayagf", "version"}
	os.Exit(m.Run())
}

func TestProgram(t *testing.T) {
	main()
}
