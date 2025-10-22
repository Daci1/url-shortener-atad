package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PingHandler struct{}

type PingAttributes struct {
	Message string `json:"message"`
}

type PingData struct {
	Type       string         `json:"type"`
	Attributes PingAttributes `json:"attributes"`
}

type PingResponse struct {
	Data PingData `json:"data"`
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, PingResponse{
		Data: PingData{
			Type:       "pings",
			Attributes: PingAttributes{Message: "pong"},
		},
	})
}
