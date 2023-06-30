package endpoint

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DayLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DayLeft()
	s := fmt.Sprintf("Day left: %d", d)
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return errors.New("Error CTX Server")
	}
	return nil
}
