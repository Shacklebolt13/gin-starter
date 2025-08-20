package controller

import (
	"cert/internal/controller/health"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UrlMapping interface {
	MapAll(engine *gin.Engine)
}

type urlMapping struct {
	health health.HealthController
}

func (u *urlMapping) MapAll(engine *gin.Engine) {
	log.Info().Msg("Building Health Mappings")
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
