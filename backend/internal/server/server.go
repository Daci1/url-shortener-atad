package server

import (
	"github.com/Daci1/url-shortener-atad/internal/server/handler"
	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()

	apiV1 := e.Group("/api/v1")

	pingHandler := handler.NewPingHandler()
	urlHandler := handler.NewUrlHandler()

	apiV1.GET("/ping", pingHandler.Ping)
	apiV1.GET("/urls/:url", urlHandler.RedirectUrl)
	apiV1.POST("/urls", urlHandler.GenerateShortenedUrl)

	return e
}
