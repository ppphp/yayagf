package spec

import (
	"gitlab.papegames.com/fengche/yayagf/pkg/spec/it"
	"testing"
)

func TestRecoverOrError(t *testing.T) {
	it.Should("panic").Run(func() {
		if err := RecoverOrError(func() {
			panic("here")
		}); err != nil {
			t.Errorf("not panic, %v", err)
		}
		if err := RecoverOrError(func() {
		}); err == nil {
			t.Errorf("panicked, %v", err)
		}
	})
}

func TestGolden(t *testing.T) {
	it.Should("new").Run(func() {
		g := LoadGolden("testdata/spec/golden.golden")
		if g == nil {
			t.Errorf("g is nil")
		}
		h := LoadGolden("testdata/spec/null.golden")
		if h == nil {
			t.Errorf("h is nil")
		}
	})
}

func TestGolden_Compare(t *testing.T) {
	it.Should("true").Run(func() {
		g := LoadGolden("testdata/golden.golden")
		if g == nil {
			t.Errorf("g is nil")
		}
		if g.Compare([]byte{}) {
			t.Errorf("g is not empty")
		}
	})
	it.Should("false").Run(func() {
		h := LoadGolden("testdata/null.golden")
		if h == nil {
			t.Errorf("h is nil")
		}
		if !h.Compare([]byte{}) {
			t.Errorf("h is not empty")
		}
	})
}

func TestGolden_Update(t *testing.T) {
	it.Should("ok").Run(func() {
		g := LoadGolden("testdata/golden.golden")
		if g.Compare([]byte("a")) {
			t.Errorf("g is not empty")
		}
		if err := g.Update(); err != nil {
			t.Errorf("g update error %v", err)
		}
	})
}
