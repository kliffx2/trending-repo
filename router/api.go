package router

import (
	"github.com/kliffx2/trending-repo/handler"
	middleware "github.com/kliffx2/trending-repo/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter()  {
	//user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	
	//profile
	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
	user.PUT("/profile/update", api.UserHandler.UpdateProfile)
}