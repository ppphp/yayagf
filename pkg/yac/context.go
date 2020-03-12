package yac

import (
	"fmt"
	"math/rand"
	"net"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	eventID string
	Logger  *logrus.Logger
}

var Logger = logrus.New()

func FromGin(c *gin.Context) *Context {
	ctx := &Context{Context: c, eventID: fmt.Sprint(rand.Int63()), Logger: Logger}
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

	realIP := strings.TrimSpace(c.Request.Header.Get("X-Real-Ip"))
	if realIP != "" {
		return realIP
	}
	remoteIP := strings.Split(c.Request.RemoteAddr, ":")
	if len(remoteIP) > 0 {
		return remoteIP[0]
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
