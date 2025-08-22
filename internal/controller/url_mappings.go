package controller

import (
	"gin-starter/internal/controller/health"
	"gin-starter/internal/controller/user"

	"github.com/gin-gonic/gin"
)

type UrlMapping interface {
	MapAll(engine *gin.Engine)
}

type urlMapping struct {
	health health.HealthController
	user   user.UserController
}

func (u *urlMapping) MapAll(engine *gin.Engine) {
	healthGroup := engine.Group("/health")
	u.health.Map(healthGroup)

	userGroup := engine.Group("/users")
	u.user.Map(userGroup)
}

func NewUrlMapping(
	health health.HealthController,
	user user.UserController,
) UrlMapping {
	return &urlMapping{
		health: health,
		user:   user,
	}
}
