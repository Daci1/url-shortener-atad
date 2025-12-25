package server

import (
	"github.com/Daci1/url-shortener-atad/internal/security"
	"github.com/Daci1/url-shortener-atad/internal/server/handler"
	"github.com/Daci1/url-shortener-atad/internal/server/middleware"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	e := echo.New()
	e.Use(echoMiddleware.CORS())

	apiV1 := e.Group("/api/v1")

	pingHandler := handler.NewPingHandler()
	userHandler := handler.NewUserHandler()
	urlHandler := handler.NewUrlHandler()

	// Public endpoints
	apiV1.POST("/users", userHandler.RegisterUser)
	apiV1.POST("/users/login", userHandler.LoginUser)
	apiV1.GET("/ping", pingHandler.Ping)
	apiV1.POST("/urls", urlHandler.GenerateShortenedUrl)
	apiV1.GET("/urls/:url", urlHandler.RedirectUrl)

	// Authenticated endpoints
	authGroupV1 := apiV1.Group("")
	authGroupV1.Use(echojwt.JWT(security.AccessTokenSecret))
	authGroupV1.POST("/urls/users/:user", urlHandler.CreateUrlForUser, middleware.UserMatchesToken)
	authGroupV1.POST("/urls/users/:user/custom", urlHandler.CreateCustomUrl, middleware.UserMatchesToken)

	// TODO: add refresh token endpoint

	return e
}
