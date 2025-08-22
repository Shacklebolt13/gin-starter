package main

import (
	"gin-starter/di"
	"gin-starter/internal/utils/errs"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	appConfig := errs.Fatal(di.ProvideAppConfig())
	di.ProvideUrlMappings().MapAll(engine)
	gin.SetMode(appConfig.Env.GIN_MODE)

	_ = engine.Run()
}
