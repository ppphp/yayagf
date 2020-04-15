package swag

import (
	"gitlab.papegames.com/fengche/yayagf/pkg/spec/it"
	"testing"
)

func TestNew(t *testing.T) {
	it.Should("new").Run(func() {
		p := New()
		if p == nil {
			t.Errorf("new nil")
		}
	})
}
