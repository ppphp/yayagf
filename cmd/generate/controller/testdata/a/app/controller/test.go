package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.papegames.com/fengche/yayagf/pkg/maotai"
)

// IndexTEST godoc
// @Summary TEST
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {int} int 0
// @Failure 200 {int} int 0
// @Failure 200 {int} int 0
// @Router /test [get]
func IndexTEST(c *maotai.Context) (int, string, gin.H) {
	return 0, "", nil
}

// CreateTEST godoc
// @Summary CreateTEST
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {int} int 0
// @Failure 200 {int} int 0
// @Failure 200 {int} int 0
// @Router /test [post]
func CreateTEST(c *maotai.Context) (int, string, gin.H) {
	return 0, "", nil
}
