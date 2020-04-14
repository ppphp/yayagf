package main

import (
	"os"
	"testing"
)

func TestMain1(t *testing.T) {
	os.Args = []string{"yayagf", "version"}
	main()
}
