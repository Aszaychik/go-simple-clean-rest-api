package main

import (
	"Aszaychik/go-simple-clean-rest-api/config"
	"Aszaychik/go-simple-clean-rest-api/controller"
	"Aszaychik/go-simple-clean-rest-api/model/domain"
	"Aszaychik/go-simple-clean-rest-api/repository"
	"Aszaychik/go-simple-clean-rest-api/routes"
	"Aszaychik/go-simple-clean-rest-api/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	validate := validator.New()

	db, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.User{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))
	
	routes.NewUserRoutes(e, userController)

	e.Logger.Fatal(e.Start(":8080"))
}