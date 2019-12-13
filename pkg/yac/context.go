package yac

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	eventID string
}

func FromGin(c *gin.Context) *Context {
	ctx := &Context{Context: c}
	return ctx
}

func (c *Context) ClientIP() string {
	clientIP := c.Request.Header.Get("X-Forwarded-For")
	for _, ip := range strings.Split(clientIP, ",") {
		if ip != "" {
			if isPublicIP(net.ParseIP(ip)) {
				return strings.TrimSpace(ip)
			}
		}
	}
	if clientIP == "" {
		return strings.TrimSpace(c.Request.Header.Get("X-Real-Ip"))
	}
	return ""
}

func isPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
