package middleware

import (
	"github.com/kliffx2/trending-repo/model"
	"github.com/kliffx2/trending-repo/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.JWT_KEY),
	}
	return middleware.JWTWithConfig(config)
}