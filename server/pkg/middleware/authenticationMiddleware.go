package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/faysal146/golang-vue-http-crud-app/server/database"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/helpers"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func AuthenticationMiddleware(c *fiber.Ctx) error {
	// check it's content type json or not
	authorizationHeader := c.GetReqHeaders()["Authorization"]
	// check authorization header

	if authorizationHeader == "" {
		return fiber.NewError(http.StatusUnauthorized, "token not found")
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	if strings.HasPrefix(authorizationHeader, "Bearer ") && helpers.IsTokenValid(tokenString) {
		// token format valid
		tokenPayload, tokenErr := helpers.VerifyToken(tokenString)
		if tokenErr != nil {
			return fiber.NewError(http.StatusUnauthorized, "unauthorized")
		} else {
			var userData model.User
			// check if user exist in the database
			err := database.DBClient.Model(&model.User{}).Where("id = ?", tokenPayload.Id).Find(&userData).Error
			if err == nil && userData.ID == tokenPayload.Id {
				// check is user change password after token generated
				if userData.PasswordUpdatedAt.Add(-time.Microsecond*100).Local().Unix() > tokenPayload.IssuedAt {
					// password change
					return fiber.NewError(http.StatusUnauthorized, "password change, please login again")
				} else {
					// everything is ok
					c.Locals("UserData", userData)
					fmt.Println("user data", userData)
					return c.Next()
				}
			} else {
				return fiber.NewError(http.StatusUnauthorized, "user not found")
			}
		}
	} else {
		return fiber.NewError(http.StatusUnauthorized, "invalid token / token not found")
	}
}
