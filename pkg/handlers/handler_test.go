package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMountALotOfThingToEndpoint(t *testing.T) {
	g := gin.Default()
	MountALotOfThingToEndpoint(g.Group("admin"), WithMetric(), WithSwagger(""))
}
