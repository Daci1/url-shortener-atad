package handler

import (
	"fmt"
	"net/http"

	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/server/response"
	"github.com/Daci1/url-shortener-atad/internal/service"
	"github.com/Daci1/url-shortener-atad/internal/shortener"
	"github.com/labstack/echo/v4"
)

type UrlHandler struct {
	s *service.UrlService
}

func NewUrlHandler() *UrlHandler {
	urlService := service.NewUrlService(db.NewUrlRepository(db.GetDB()))
	return &UrlHandler{
		s: urlService,
	}
}

func (h *UrlHandler) RedirectUrl(c echo.Context) error {
	shortUrl := c.Param("url")
	if len(shortUrl) != shortener.MaxUrlLength {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid short url"))
	}

	originalUrl, err := h.s.GetUrl(shortUrl)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerErrorResponse())
	}

	if originalUrl == "" {
		return c.JSON(http.StatusNotFound, response.NewErrorResponse(http.StatusNotFound, "Short url not found"))
	}

	return c.Redirect(http.StatusMovedPermanently, originalUrl)
}

func (h *UrlHandler) GenerateShortenedUrl(c echo.Context) error {
	var req response.ApiRequest[response.CreateUrlRequestAttributes]
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid request body"))
	}

	res, err := h.s.CreateUrl(req.Data.Attributes.OriginalUrl)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusCreated, res)
}
