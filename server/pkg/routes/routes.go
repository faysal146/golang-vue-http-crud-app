package routes

import (
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/controllers"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(r fiber.Router) {
	// check it's content type json or not
	r.Use(func(c *fiber.Ctx) error {
		if c.GetReqHeaders()["Content-Type"] != "application/json" {
			return fiber.NewError(fiber.StatusUnsupportedMediaType, "Content Type is not application/json")
		} else {
			return c.Next()
		}
	})
	authRoute := r.Group("/auth")
	authRoute.Post("/login", controllers.LoginUser)
	authRoute.Post("/register", controllers.RegisterUser)
	authRoute.Post("/refresh-token", middleware.AuthenticationMiddleware, controllers.RefreshToken)
}
