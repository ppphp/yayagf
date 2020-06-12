// 利好茅台，奥利给！
package maotai

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.papegames.com/fengche/yayagf/pkg/prom"
	"net/http"
	"time"
)

type MaoTai struct {
	*gin.Engine
	UrlConn *prometheus.GaugeVec
	TTLHist *prometheus.HistogramVec
}

func (m *MaoTai) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.Use(middleware...)
}

func (m *MaoTai) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.Handle(httpMethod, relativePath, handlers...)
}

func (m *MaoTai) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.GET(relativePath, handlers...)
}

func (m *MaoTai) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.POST(relativePath, handlers...)
}

func (m *MaoTai) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.PUT(relativePath, handlers...)
}

func (m *MaoTai) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.DELETE(relativePath, handlers...)
}

func (m *MaoTai) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.PATCH(relativePath, handlers...)
}

func (m *MaoTai) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.OPTIONS(relativePath, handlers...)
}

func (m *MaoTai) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return m.Engine.HEAD(relativePath, handlers...)
}

func New(project string) *MaoTai {
	m := &MaoTai{}

	m.UrlConn = prom.UrlConnection(project)
	m.TTLHist = prom.UrlTTL(project)

	m.Engine = gin.New()
	return m
}

func Default(project string) *MaoTai {
	m := &MaoTai{}

	m.UrlConn = prom.UrlConnection(project)
	m.TTLHist = prom.UrlTTL(project)

	m.Engine = gin.Default()
	return m
}

func NikkiSerializer(m *MaoTai, controller func(*gin.Context) (int, string, gin.H)) func(*gin.Context) {
	return func(c *gin.Context) {
		var ret int
		var msg string
		mp, mret := map[string]interface{}{}, map[string]interface{}{}
		m.UrlConn.WithLabelValues(c.Request.URL.Path, c.Request.Method).Add(1)
		defer func(t time.Time) {
			m.TTLHist.WithLabelValues(c.Request.URL.Path, c.Request.Method, fmt.Sprint(ret)).Observe(time.Since(t).Seconds())
			m.UrlConn.WithLabelValues(c.Request.URL.Path, c.Request.Method).Add(-1)
		}(time.Now())
		ret, msg, mp = controller(c)
		for k, v := range mp {
			mret[k] = v
		}
		mret["ret"] = ret
		mret["msg"] = msg
		mret["timestamp"] = time.Now().Unix()
		c.JSON(http.StatusOK, mret)
	}
}

func TDSSerializer(m *MaoTai, controller func(*gin.Context) (int, string, gin.H)) func(*gin.Context) {
	return func(c *gin.Context) {
		var ret int
		var msg string
		mp, mret := map[string]interface{}{}, map[string]interface{}{}
		m.UrlConn.WithLabelValues(c.Request.URL.Path, c.Request.Method).Add(1)
		defer func(t time.Time) {
			m.TTLHist.WithLabelValues(c.Request.URL.Path, c.Request.Method, fmt.Sprint(ret)).Observe(time.Since(t).Seconds())
			m.UrlConn.WithLabelValues(c.Request.URL.Path, c.Request.Method).Add(-1)
		}(time.Now())
		ret, msg, mp = controller(c)
		for k, v := range mp {
			mret[k] = v
		}
		mret["iRet"] = ret
		mret["sMsg"] = msg
		mret["timestamp"] = time.Now().Unix()
		c.JSON(http.StatusOK, mret)
	}
}

func PlainSerializer(m *MaoTai, controller func(*gin.Context) (int, string, interface{})) func(*gin.Context) {
	return func(c *gin.Context) {
		var ret int
		var mp interface{}
		m.UrlConn.WithLabelValues(c.Request.URL.Path, c.Request.Method).Add(1)
		defer func(t time.Time) {
			m.TTLHist.WithLabelValues(c.Request.URL.Path, c.Request.Method, fmt.Sprint(ret)).Observe(time.Since(t).Seconds())
			m.UrlConn.WithLabelValues(c.Request.URL.Path, c.Request.Method).Add(-1)
		}(time.Now())
		ret, _, mp = controller(c)
		c.JSON(http.StatusOK, mp)
	}
}
