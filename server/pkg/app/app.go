package app

import (
	"fmt"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/helpers"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func Run(addr string) {
	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.DefaultErrorHandler,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(app.Stack())
	})
	router := app.Group("/api/v1")
	routes.InitializeRoutes(router)

	fmt.Println("server running on port ", addr)
	if err := app.Listen(addr); err != nil {
		fmt.Println("could not start server: ")
		panic(err)
	}
}
