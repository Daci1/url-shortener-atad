package server

import (
	"github.com/Daci1/url-shortener-atad/internal/server/handler"
	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()

	apiV1 := e.Group("/api/v1")

	pingHandler := handler.NewPingHandler()

	apiV1.GET("/ping", pingHandler.Ping)

	return e
}
