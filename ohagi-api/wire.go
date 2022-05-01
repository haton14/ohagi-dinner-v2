//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/haton14/ohagi-dinner/ohagi-api/handler"
	"github.com/haton14/ohagi-dinner/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

func InitializeApp(e *echo.Echo) App {
	wire.Build(NewApp, handler.NewAuth, usecase.NewAuth)
	return App{}
}
