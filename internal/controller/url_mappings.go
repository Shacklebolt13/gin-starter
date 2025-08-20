package controller

import (
	"cert/internal/controller/health"
	"cert/internal/utils"

	"github.com/gin-gonic/gin"
)

type UrlMapping interface {
	MapAll(engine *gin.Engine)
}

type urlMapping struct {
	health health.HealthController
}

func (u *urlMapping) MapAll(engine *gin.Engine) {
	engine.Group("/health", utils.Fatal(u.health.Map(engine)))
}

func NewUrlMapping(
	health health.HealthController,
) UrlMapping {
	return &urlMapping{
		health: health,
	}
}
