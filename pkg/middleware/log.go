package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 简单的config，复杂的做不来

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func entries() *logrus.Entry {
	fun, filename, line, _ := runtime.Caller(2)
	f := runtime.CallersFrames([]uintptr{fun})
	m, _ := f.Next()
	return logger.WithFields(logrus.Fields{
		"position": fmt.Sprintf("%v:%v", filename, line),
		"function": m.Function,
	})
}

func Debugf(format string, args ...interface{}) {
	entries().Debugf(format, args...)
}

func Print(format string, args ...interface{}) {
	entries().Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	entries().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	entries().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	entries().Fatalf(format, args...)
}

func Infof(format string, args ...interface{}) {
	entries().Infof(format, args...)
}

type Config struct {
	FileFunc bool
	Level    int
	Format   string
	Stdout   bool
	File     bool
}

func GetLogger() *logrus.Logger {
	return logger
}

func Tweak(config Config) {
	Infof("%+v", config)
	l := logrus.New()

	// ensure level
	isValidLevel := false
	for _, level := range logrus.AllLevels {
		if level == logrus.Level(config.Level) {
			Infof("matched level %v", level)
			l.SetLevel(logrus.Level(config.Level))
			isValidLevel = true
			break
		}
	}
	if !isValidLevel {
		Warnf("not valid level %v", config.Level)
		l.SetLevel(logrus.DebugLevel)
	}

	if config.Format == "json" {
		l.SetFormatter(&logrus.JSONFormatter{})
	}
	l.SetReportCaller(config.FileFunc)

	if config.File {
		file, err := os.OpenFile("jifenbackend.log", os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logger.Warnf("open jifenbackend.log failed(%v)\n", err)
		} else {
			l.SetOutput(file)
		}
	}
	logger = l
}

func Ginrus(logger loggerEntryWithFields) gin.HandlerFunc {
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

type loggerEntryWithFields interface {
	WithFields(fields logrus.Fields) *logrus.Entry
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
