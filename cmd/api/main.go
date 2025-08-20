package main

import (
	"cert/di"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	di.ProvideUrlMappings().MapAll(router)
	_ = router.Run()
}
