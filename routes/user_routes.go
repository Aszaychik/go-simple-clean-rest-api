package routes

import (
	"Aszaychik/go-simple-clean-rest-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", userController.RegisterUserController)
	usersGroup.POST("/login", userController.LoginUserController)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.PUT("/:id", userController.UpdateUserController)
	usersGroup.GET("/:id", userController.GetUserController)
	usersGroup.GET("", userController.GetUsersController)
	usersGroup.DELETE("/:id", userController.DeleteUserController)
}