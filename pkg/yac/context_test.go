package yac

import (
	"gitlab.papegames.com/fengche/yayagf/pkg/spec/it"
	"net"
	"testing"
)

func TestIsPublicIP(t *testing.T) {
	it.Should("return local").Run(func() {
		if isPublicIP(net.ParseIP("127.0.0.1")) {
			t.Errorf("ip is not private %v", "127.0.0.1")
		}
	})
}

