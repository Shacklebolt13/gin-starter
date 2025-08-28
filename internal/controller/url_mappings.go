package controller

import (
	_ "gin-starter/docs"
	"gin-starter/internal/controller/health"
	"gin-starter/internal/controller/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type UrlMapping interface {
	MapAll(engine *gin.Engine)
}

type urlMapping struct {
	health health.HealthController
	user   user.UserController
}

func (u *urlMapping) MapAll(engine *gin.Engine) {
	metaGroup := engine.Group("/meta")
	metaGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	metaGroup.StaticFile("/swagger.json", "./docs/swagger.json")
	metaGroup.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

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
