package controllers

import (
	"book-store/pkg/domain"
	"book-store/pkg/models"
	"book-store/pkg/types"
	"net/http"

	"github.com/labstack/echo/v4"
)


type AuthController struct {
	authSvc domain.IAuthService
}

func NewAuthController(authSvc domain.IAuthService) AuthController {
	return AuthController{
		authSvc: authSvc,
	}
}

func (ac *AuthController) Signup(e echo.Context) error {
	reqUser := &types.UserRequest{}
	if err := e.Bind(reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	if err := reqUser.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	user := &models.User{
		UserName: reqUser.UserName,
		Email:    reqUser.Email,
		Password: reqUser.Password,
		Phone:    reqUser.Phone,
		Address:  reqUser.Address,
	}

	if err := ac.authSvc.Signup(user); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "user signup successful")
}

func (ac *AuthController) Signin(e echo.Context) error {
	reqCreds := &types.Credentials{}
	if err := e.Bind(reqCreds); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	
	if err := reqCreds.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := ac.authSvc.Signin(reqCreds.Username, reqCreds.Password)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	return e.JSON(http.StatusOK, map[string]string{"token": token})
}
