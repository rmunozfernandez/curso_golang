package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	/*
		go mod vendor
		Guarda todas las dependencias de los modulos ocupados.

		go mod tidy
		Limpia las dependecias no utilizadas
	*/
}
