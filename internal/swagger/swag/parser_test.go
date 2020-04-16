package swag

import (
	"gitlab.papegames.com/fengche/yayagf/pkg/spec/it"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	it.Should("new").Run(func() {
		p := New()
		if p == nil {
			t.Errorf("new nil")
		}
	})
	it.Should("option").Run(func() {
		p := New(SetLogger(os.Stdout))
		if p == nil {
			t.Errorf("new nil")
		}
	})
}
