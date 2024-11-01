package auth

import (
	"rest/controller/auth/request"
	"rest/controller/auth/response"
	"rest/controller/base"
	"rest/service/auth"

	"github.com/labstack/echo/v4"
)

func NewAuthController(as auth.AuthInterface) *AuthController {
	return &AuthController{
		authService: as,
	}
}

type AuthController struct {
	authService auth.AuthInterface
}

func (userController AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := userController.authService.Login(userLogin.ToEntities())
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	return base.SuccesResponse(c, response.FromEntities(user))
}

func RegisterHandler(c echo.Context) error {
	return base.SuccesResponse(c, nil)
}
