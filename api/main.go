package main

import (
	"context"
	"fmt"
	"github.com/echo-stream/api/application/usecase"
	"github.com/echo-stream/api/domain/validate"
	"github.com/echo-stream/api/infrastructure"
	"github.com/echo-stream/api/infrastructure/repository"
	"github.com/echo-stream/api/presentation/controller"
	"github.com/echo-stream/database/ent"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Validator = &validate.Validate{Validator: validator.New()}

	dbDriver, dbConnectionString := infrastructure.NewDatabase()
	client, err := ent.Open(dbDriver, dbConnectionString)
	if err != nil {
		fmt.Printf("failed opening connection to mysql: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)
	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		fmt.Printf("failed creating schema resources: %v", err)
	}

	sr := repository.NewSlackRepository(client)
	su := usecase.NewSlackUsecase(*sr)
	sc := controller.SlackController{
		SlackUsecase: *su,
	}
	e.POST("/sent", sc.Create)

	// 서버 시작
	e.Logger.Fatal(e.Start(":1323"))
}
