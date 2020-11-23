package log

import (
	"fmt"
	"os"
	"runtime"

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

func Infof(format string, args ...interface{}) {
	entries().Infof(format, args...)
}

type Config struct {
	Level    int
	Format   string
	Stdout   bool
	FileName string
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

	if config.FileName != "" {
		file, err := NewRotationWriter(config.FileName, Hour)
		if err != nil {
			logger.Warnf("open %v failed(%v)\n", config.FileName, err)
		} else {
			l.SetOutput(file)
		}
	}
	logger = l
}
