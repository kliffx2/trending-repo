package main

import (
	"fmt"
	"time"

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

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	go scheduleUpdateTrending(10*time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler)  {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for{
			select{
			case <- ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
