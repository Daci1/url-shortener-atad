package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserMatchesToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenAny := c.Get("user")
		if tokenAny == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		token, ok := tokenAny.(*jwt.Token)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "userId not in token")
		}

		userIdFromURL := c.Param("user")
		if userIdFromURL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing user id in URL")
		}

		if sub != userIdFromURL {
			return echo.NewHTTPError(http.StatusForbidden, "you are not allowed to access this resource")
		}

		return next(c)
	}
}
