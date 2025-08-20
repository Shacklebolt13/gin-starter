package health

import (
	"github.com/gin-gonic/gin"
)

type healthControllerImpl struct {
	HealthController
}

func (h *healthControllerImpl) Health(c *gin.Context) {
	c.Status(200)
}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}
