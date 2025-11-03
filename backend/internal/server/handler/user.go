package handler

import (
	"fmt"
	"net/http"

	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/server/response"
	"github.com/Daci1/url-shortener-atad/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s *service.UserService
}

func NewUserHandler() *UserHandler {
	userService := service.NewUserService(db.NewUserRepository(db.GetDB()))
	return &UserHandler{
		s: userService,
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req response.ApiRequest[response.RegisterRequestAttributes]
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return c.JSON(response.NewErrorResponse(http.StatusBadRequest, "Invalid request body"))
	}

	res, err := h.s.RegisterUser(req.Data.Attributes)

	if err != nil {
		fmt.Println(err)
		return c.JSON(response.NewErrorFromCustomError(err))
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	var req response.ApiRequest[response.LoginRequestAttributes]
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return c.JSON(response.NewErrorResponse(http.StatusBadRequest, "Invalid request body"))
	}

	// TODO: add error handling for register user
	res, err := h.s.LoginUser(req.Data.Attributes)

	if err != nil {
		fmt.Println(err)
		return c.JSON(response.NewErrorFromCustomError(err))
	}

	return c.JSON(http.StatusCreated, res)
}
