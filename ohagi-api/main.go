package main

import (
	"github.com/haton14/ohagi-dinner/ohagi-api/domain"
	"github.com/haton14/ohagi-dinner/ohagi-api/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := InitializeApp(echo.New(), domain.NewPetTokenMaker("", ""))
	app.Run()
}

type App struct {
	e           *echo.Echo
	authHandler handler.Auth
}

func NewApp(e *echo.Echo, authHandler handler.Auth) App {
	return App{
		e:           e,
		authHandler: authHandler,
	}
}

func (a App) Run() {
	a.authHandler.Handle()
	a.e.Logger.Fatal(a.e.Start(":8000"))
}
