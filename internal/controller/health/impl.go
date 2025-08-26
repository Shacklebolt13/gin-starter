package health

import (
	"github.com/gin-gonic/gin"
)

type healthControllerImpl struct {
	HealthController
}

// Health godoc
// @Summary Checks if the service is healthy
// @Description Returns a 200 OK status if the service is healthy
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get].
func (h *healthControllerImpl) Health(c *gin.Context) {
	c.Status(200)
}

func (h *healthControllerImpl) Map(e *gin.RouterGroup) {
	e.GET("/", h.Health)
}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}
