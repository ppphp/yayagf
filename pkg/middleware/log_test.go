package middleware

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGinrus(t *testing.T) {
	Ginrus(logrus.New())
}
