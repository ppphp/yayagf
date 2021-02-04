package controller 

// IndexTest godoc
// @Summary Test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {int} int 0
// @Failure 200 {int} int 0
// @Failure 200 {int} int 0
// @Router /test [get]
func IndexTest(c *maotai.Context) (int, string, gin.H) {
	return 0, "", nil
}

// CreateTest godoc
// @Summary CreateTest
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {int} int 0
// @Failure 200 {int} int 0
// @Failure 200 {int} int 0
// @Router /test [post]
func CreateTest(c *maotai.Context) (int, string, gin.H) {
	return 0, "", nil
}
