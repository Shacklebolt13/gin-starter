package generic

import "github.com/gin-gonic/gin"

type Controller interface {
	Map(e *gin.RouterGroup)
}
