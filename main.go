package main

import (
	"net/http"

	"github.com/kliffx2/trending-repo/db"
	"github.com/kliffx2/trending-repo/handler"
	"github.com/kliffx2/trending-repo/repository/repo_impl"
	"github.com/kliffx2/trending-repo/router"
	"github.com/labstack/echo/v4"
)

func main() {

	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "postgres",
		Password: "postgres",
		DbName: "trending_repo",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

