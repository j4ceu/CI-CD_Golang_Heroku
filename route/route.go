package route

import (
	"CICD_GolangtoHeroku/configs"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.POST("/login", configs.UserController.LoginUser)
	e.POST("/users", configs.UserController.CreateUser)

	auth := e.Group("/auth", middleware.JWT([]byte("jace")))

	auth.GET("/users", configs.UserController.GetAllUsers)

	e.Logger.Fatal(e.Start(":" + port))

	return e

}
