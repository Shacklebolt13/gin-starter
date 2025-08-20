package controller

import (
	"gin-starter/internal/controller/health"

	"github.com/gin-gonic/gin"
)

type UrlMapping interface {
	MapAll(engine *gin.Engine)
}

type urlMapping struct {
	health health.HealthController
}

func (u *urlMapping) MapAll(engine *gin.Engine) {
	healthGroup := engine.Group("/health")
	u.health.Map(healthGroup)
}

func NewUrlMapping(
	health health.HealthController,
) UrlMapping {
	return &urlMapping{
		health: health,
	}
}
