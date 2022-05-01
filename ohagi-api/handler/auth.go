package handler

import (
	"fmt"

	"github.com/haton14/ohagi-dinner/ohagi-api/middleware"
	"github.com/haton14/ohagi-dinner/ohagi-api/presenter"
	"github.com/haton14/ohagi-dinner/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

func NewAuth(e *echo.Echo, usecase usecase.Auth) Auth {
	return &auth{
		e:       e,
		usecase: usecase,
	}
}

type Auth interface {
	Handle()
}

type auth struct {
	e       *echo.Echo
	usecase usecase.Auth
}

func (a *auth) Handle() {
	a.e.POST("/login", a.login, middleware.FirebaseAuth())
}

func (a *auth) login(c echo.Context) error {
	userID := fmt.Sprint(c.Get("firebase_user_id"))
	request := usecase.LoginRequest{UserID: userID}
	return a.usecase.GenarateToken(request, presenter.NewAuth(c))
}

type loginRequest struct {
	userID string
}
