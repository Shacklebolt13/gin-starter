package health

import (
	"gin-starter/internal/controller/generic"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	generic.Controller
	Health(c *gin.Context)
}
