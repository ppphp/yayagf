package log

import "github.com/sirupsen/logrus"

var Logger = logrus.New()

func Printf(format string, args ...interface{}) {
	Logger.Printf(format, args...)
}

func Println(args ...interface{}) {
	Logger.Println(args...)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
