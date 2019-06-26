package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()
	
	e.GET("/sample", func(c echo.Context) error {
		return c.String(http.StatusOK, "Helloworld")
	})
	return e
}