package main

import (
	"gin-starter/di"

	"github.com/gin-gonic/gin"
)

// @title Gin Starter API
// @version 1.0.0.
func main() {
	engine := gin.Default()
	di.ProvideUrlMappings().MapAll(engine)

	_ = engine.Run()
}
