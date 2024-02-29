package middleware

import (
	"net/http"

	"github.com/kliffx2/trending-repo/model"
	"github.com/kliffx2/trending-repo/model/req"
	"github.com/labstack/echo/v4"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// handle logic
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}

			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "You cannot call this API!",
					Data:       nil,
				})
			}

			return next(c)
		}
	}
}