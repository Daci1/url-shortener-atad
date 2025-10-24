package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UrlHandler struct{}

func NewUrlHandler() *UrlHandler {
	return &UrlHandler{}
}

func (h *UrlHandler) GenerateShortenedUrl(c echo.Context) error {
	const originalUrl = "https://google.com"
	return c.Redirect(http.StatusMovedPermanently, originalUrl)
}
