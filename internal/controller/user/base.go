package user

import (
	"gin-starter/internal/controller/generic"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	generic.Controller
	RegisterUserByPassword(c *gin.Context)
}
