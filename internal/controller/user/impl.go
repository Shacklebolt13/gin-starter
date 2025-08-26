package user

import (
	"gin-starter/internal/service/app/user"
	"gin-starter/internal/service/app/user/create"
	"gin-starter/internal/service/app/user/login"
	"gin-starter/internal/utils/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userControllerImpl struct {
	uSvc user.UserService
}

func (u *userControllerImpl) Map(e *gin.RouterGroup) {
	e.POST("/sign-up/password/", u.RegisterUserByPassword)
	e.POST("/login/password/", u.LoginByPassword)
}

func NewUserController(uSvc *user.UserService) UserController {
	return &userControllerImpl{
		uSvc: *uSvc,
	}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param request body create.CreateUserByPasswordRequest true "User registration request"
// @Success 201 {object} errs.Incident
// @Failure 404 {object} errs.Incident
// @Failure 400 {object} errs.Incident
// @Failure 500 {object} errs.Incident
// @Router /user/sign-up/password [post].
func (u *userControllerImpl) RegisterUserByPassword(c *gin.Context) {
	var request create.CreateUserByPasswordRequest
	if err := c.BindJSON(&request); err != nil {
		if e, ok := err.(errs.ValidationErr); ok {
			c.JSON(http.StatusBadRequest, errs.Incident{
				Type:    errs.RequestValidationError,
				Message: e.Error(),
				Detail:  e.Map(),
			})

			return
		} else {
			c.JSON(http.StatusInternalServerError, errs.Incident{
				Type:    errs.RequestValidationError,
				Message: err.Error(),
			})

			return
		}
	}

	result, err := u.uSvc.Create.RegisterUserByPassword(c.Request.Context(), request)
	if err == nil {
		c.JSON(http.StatusCreated, errs.Incident{
			Type:   errs.Success,
			Detail: result,
		})

		return
	}

	if inc, ok := err.(*errs.Incident); ok {
		c.JSON(int(inc.Type), inc)
	}
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Accept json
// @Produce json
// @Param request body login.LoginByPasswordRequest true "User login request"
// @Success 200 {object} errs.Incident
// @Failure 400 {object} errs.Incident
// @Failure 404 {object} errs.Incident
// @Failure 500 {object} errs.Incident
// @Router /user/login/password [post].
func (u *userControllerImpl) LoginByPassword(c *gin.Context) {
	var request login.LoginByPasswordRequest
	if err := c.BindJSON(&request); err != nil {
		if e, ok := err.(errs.ValidationErr); ok {
			c.JSON(http.StatusBadGateway, errs.Incident{
				Type:    errs.RequestValidationError,
				Message: e.Error(),
				Detail:  e.Map(),
			})

			return
		} else {
			c.JSON(http.StatusBadRequest, errs.Incident{
				Type:    errs.RequestValidationError,
				Message: err.Error(),
			})

			return
		}
	}

	result, err := u.uSvc.Login.LoginByPassword(c.Request.Context(), request)
	if err == nil {
		c.JSON(http.StatusOK, errs.Incident{
			Type:   errs.Success,
			Detail: result,
		})

		return
	}

	if inc, ok := err.(*errs.Incident); ok {
		c.JSON(int(inc.Type), inc)
	}
}
