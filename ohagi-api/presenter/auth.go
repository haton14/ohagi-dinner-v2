package presenter

import (
	"net/http"

	"github.com/haton14/ohagi-dinner/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

func NewAuth(c echo.Context) usecase.LoginResponse {
	return &auth{c: c}
}

type auth struct {
	c echo.Context
}

func (a auth) RenderLoginResult(petsToken string) error {
	return a.c.JSON(http.StatusOK, LoginResultResponse{PetsToken: petsToken})
}

func (a auth) RenderFailure(err error) error {
	return sendResponse(a.c, http.StatusInternalServerError)
}

type LoginResultResponse struct {
	PetsToken string `json:"pets_token"`
}
