package controllers

import (
	"net/http"
	"strings"

	"github.com/faysal146/golang-vue-http-crud-app/server/database"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/helpers"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var bodyData = new(model.RegisterBody)
	if parseErr := c.BodyParser(bodyData); parseErr != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid body")
	}
	validationErr := validator.New().Struct(bodyData)
	if validationErr != nil {
		// custom error messages customValidationErr{ email: "invalid email", .... }
		var customValidationErr = make(map[string]string)
		// range over all errors
		for _, err := range validationErr.(validator.ValidationErrors) {
			// get human friendly error message -> "email", "invalid email address"
			key, msg := helpers.GetValidationMessage(err)
			customValidationErr[key] = msg
		}
		// response back with error message
		return c.Status(http.StatusBadRequest).JSON(customValidationErr)
	} else {
		/*
			 	* check if user already exist
					- by email
					- by username
				if user not exist save user data to the database create jwt and refresh token
		*/

		var userExist bool
		database.DBClient.Model(&model.User{}).Where("email = ?", strings.ToLower(bodyData.Email)).Or("username = ?", bodyData.Username).Find(&userExist)

		if userExist {
			return fiber.NewError(http.StatusForbidden, "user already exist")
		} else {
			var userData = model.User{
				FirstName: bodyData.FirstName,
				LastName:  bodyData.LastName,
				Username:  bodyData.Username,
				Email:     bodyData.Email,
				Password:  bodyData.Password,
			}
			// TODO: remove password field from response
			// respField := []string{"id", "email", "first_name", "last_name", "username", "created_at", "updated_at"}
			database.DBClient.Model(&model.User{}).Create(&userData)
			authToken, refreshToken := helpers.GenerateAuthToken(userData.ID, userData.Email)
			userData.Token = authToken
			userData.RefreshToken = refreshToken
			database.DBClient.Model(&userData).Updates(model.User{Token: authToken, RefreshToken: refreshToken}).Where("id = ?", userData.ID)
			return c.Status(fiber.StatusCreated).JSON(userData)
		}
	}
}

func LoginUser(c *fiber.Ctx) error {
	// login either username or email and password

	var loginBody = new(model.LoginBody)
	if parseErr := c.BodyParser(loginBody); parseErr != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid body")
	}
	// validate email and password
	validationErr := validator.New().Struct(loginBody)
	if validationErr != nil {
		// custom error messages customValidationErr{ email: "invalid email", .... }
		var customValidationErr = make(map[string]string)
		// range over all errors
		for _, err := range validationErr.(validator.ValidationErrors) {
			// get human friendly error message -> "email", "invalid email address"
			key, msg := helpers.GetValidationMessage(err)
			customValidationErr[key] = msg
		}
		// response back with error message
		return c.Status(http.StatusBadRequest).JSON(customValidationErr)
	} else {
		var userdata model.User
		// file format correct

		// check is user exist
		database.DBClient.Model(&model.User{}).Where("email = ?", strings.ToLower(loginBody.Email)).Find(&userdata)
		if userdata.Email == strings.ToLower(loginBody.Email) && strings.TrimSpace(userdata.ID) != "" {
			// user found

			// check password
			if userdata.VerifyPassword(loginBody.Password) {
				// password match
				authToken, refreshToken := helpers.GenerateAuthToken(userdata.ID, userdata.Email)
				// update token on database
				err := database.DBClient.Model(&userdata).Updates(model.User{Token: authToken, RefreshToken: refreshToken}).Where("id = ?", userdata.ID).Error
				if err != nil {
					return fiber.NewError(http.StatusInternalServerError, "internal server error")
				}
				userdata.Token = authToken
				userdata.RefreshToken = refreshToken
				return c.Status(fiber.StatusOK).JSON(userdata)
			} else {
				// password don't match
				return fiber.NewError(http.StatusUnauthorized, "invalid password")
			}
		} else {
			// user not found
			return fiber.NewError(http.StatusForbidden, "user not found")
		}
	}
}

func RefreshToken(c *fiber.Ctx) error {
	var tokenFormBody = struct {
		Token string `json:"token"`
	}{}
	if parseErr := c.BodyParser(&tokenFormBody); parseErr != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid req body. token is required")
	}
	if helpers.IsTokenValid(tokenFormBody.Token) {
		if err := helpers.VerifyRefreshToken(tokenFormBody.Token); err != nil {
			return fiber.NewError(http.StatusBadRequest, "invalid token")
		} else {
			userdata := c.Locals("UserData").(model.User)
			authToken, refreshToken := helpers.GenerateAuthToken(userdata.ID, userdata.Email)
			return c.Status(fiber.StatusOK).JSON(map[string]string{"token": authToken, "refresh_token": refreshToken})
		}
	} else {
		return fiber.NewError(http.StatusBadRequest, "invalid token")
	}
}
