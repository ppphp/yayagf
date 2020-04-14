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
