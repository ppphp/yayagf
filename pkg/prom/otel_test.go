package prom

import (
	"testing"

	"go.opentelemetry.io/otel/metric/global"
)

func TestOtelDiskUsage(t *testing.T) {
	OtelDiskUsage(global.Meter("test"))
}
