package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Ginrus(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		req := ""
		if !c.IsWebsocket() {
			data, err := c.GetRawData()
			if err != nil {
				_ = c.Error(err)
			} else {
				req = string(data)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		resp := blw.body.String()
		end := time.Now()
		latency := end.Sub(start)

		entry := logger.WithFields(logrus.Fields{
			"status":       c.Writer.Status(),
			"method":       c.Request.Method,
			"request":      req,
			"response":     resp,
			"path":         path,
			"query":        query,
			"ip":           c.ClientIP(),
			"latency":      latency,
			"user-agent":   c.Request.UserAgent(),
			"time":         end.Format(time.RFC3339),
			"web-socket":   c.IsWebsocket(),
			"aborted":      c.IsAborted(),
			"content-type": c.ContentType(),
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
