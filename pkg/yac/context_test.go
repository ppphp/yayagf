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
	it.Should("return internal").
		Run(func() {
			if isPublicIP(net.ParseIP("127.0.0.1")) {
				t.Errorf("ip is not private %v", "127.0.0.1")
			}
		}).
		Run(func() {
			if isPublicIP(net.ParseIP("10.0.0.1")) {
				t.Errorf("ip is not private %v", "127.0.0.1")
			}
		}).
		Run(func() {
			if isPublicIP(net.ParseIP("172.16.0.1")) {
				t.Errorf("ip is not private %v", "127.0.0.1")
			}
		}).
		Run(func() {
			if isPublicIP(net.ParseIP("192.168.0.1")) {
				t.Errorf("ip is not private %v", "127.0.0.1")
			}
		}).
		Run(func() {
			if !isPublicIP(net.ParseIP("1.1.1.1")) {
				t.Errorf("ip is not private %v", "127.0.0.1")
			}
		}).
		Run(func() {
			if isPublicIP(net.ParseIP("1.1.1.1.?")) {
				t.Errorf("ip is not private %v", "127.0.0.1")
			}
		})
}
