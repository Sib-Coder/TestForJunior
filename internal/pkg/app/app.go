package app

import (
	"TestForJunior/internal/app/endpoint"
	"TestForJunior/internal/app/mw"
	"TestForJunior/internal/app/service"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server Runnig")

	err := a.echo.Start(":8090")
	if err != nil {
		log.Println(errors.New("Error Start Service"))
	}
	return nil
}