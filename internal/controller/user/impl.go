package user

import (
	"gin-starter/internal/service/app/user"
	"gin-starter/internal/service/app/user/create"
	"gin-starter/internal/utils/errs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userControllerImpl struct {
	uSvc user.UserService
}

func (u *userControllerImpl) RegisterUserByPassword(c *gin.Context) {
	var request create.CreateUserByPasswordRequest
	if err := c.BindJSON(&request); err != nil {
		errMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errMap[e.Field()] = e.Tag()
		}

		c.JSON(400, errs.Incident{
			Type:    errs.RequestValidationError,
			Message: "Invalid request format",
			Detail:  errMap,
		})
	}

	result, err := u.uSvc.Create.RegisterUserByPassword(c.Request.Context(), request)
	if err == nil {
		c.JSON(200, errs.Incident{
			Type:    errs.Success,
			Message: "User registered successfully",
			Detail:  result,
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
