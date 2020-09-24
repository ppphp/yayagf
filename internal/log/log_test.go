package log

import "testing"

func TestDebugf(t *testing.T) {
	Debugf("1")
}

func TestPrintf(t *testing.T) {
	Printf("1")
}

func TestPrintln(t *testing.T) {
	Println("1")
}

func TestErrorf(t *testing.T) {
	Errorf("1")
}
