package user

import (
	"gin-starter/internal/service/app/user"
	"gin-starter/internal/service/app/user/create"
	"gin-starter/internal/service/app/user/login"
	"gin-starter/internal/utils/errs"

	"github.com/gin-gonic/gin"
)

type userControllerImpl struct {
	uSvc user.UserService
}

func (u *userControllerImpl) RegisterUserByPassword(c *gin.Context) {
	var request create.CreateUserByPasswordRequest
	if err := c.BindJSON(&request); err != nil {
		e := err.(errs.ValidationErr)
		c.JSON(400, errs.Incident{
			Type:    errs.RequestValidationError,
			Message: e.Error(),
			Detail:  e.Map(),
		})
	}

	result, err := u.uSvc.Create.RegisterUserByPassword(c.Request.Context(), request)
	if err == nil {
		c.JSON(200, errs.Incident{
			Type:   errs.Success,
			Detail: result,
		})
	}

	if inc, ok := err.(*errs.Incident); ok {
		c.JSON(int(inc.Type), inc)
	}
}

func (u *userControllerImpl) LoginByPassword(c *gin.Context) {
	var request login.LoginByPasswordRequest
	if err := c.BindJSON(&request); err != nil {
		e := err.(errs.ValidationErr)
		c.JSON(400, errs.Incident{
			Type:    errs.RequestValidationError,
			Message: e.Error(),
			Detail:  e.Map(),
		})
	}

	result, err := u.uSvc.Login.LoginByPassword(c.Request.Context(), request)
	if err == nil {
		c.JSON(200, errs.Incident{
			Type:   errs.Success,
			Detail: result,
		})
	}

	if inc, ok := err.(*errs.Incident); ok {
		c.JSON(int(inc.Type), inc)
	}
}

func (u *userControllerImpl) Map(e *gin.RouterGroup) {
	e.GET("/sign-up/password/", u.RegisterUserByPassword)
}

func NewUserController(uSvc *user.UserService) UserController {
	return &userControllerImpl{
		uSvc: *uSvc,
	}
}
