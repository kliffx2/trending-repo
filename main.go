package main

import (
	"github.com/kliffx2/trending-repo/db"
	"github.com/kliffx2/trending-repo/handler"
	"github.com/kliffx2/trending-repo/helper"
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
	
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

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


