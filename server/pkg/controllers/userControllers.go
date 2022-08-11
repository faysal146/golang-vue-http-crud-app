package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/faysal146/golang-vue-http-crud-app/server/database"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/helpers"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/model"
	"github.com/go-playground/validator/v10"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// check it's content type json or not
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// response error
		helpers.RespondWithError(w, http.StatusUnsupportedMediaType, "Content Type is not application/json")
		return
	}
	var bodyData model.RegisterBody
	// convert req body data into json
	json.NewDecoder(r.Body).Decode(&bodyData)
	// validate req body data
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
		helpers.RespondWithError(w, http.StatusBadRequest, customValidationErr)
	} else {
		/*
			 	* check if user already exist
					- by email
					- by username
				if user not exist save user data to the database create jwt and refresh token
		*/

		var userExist bool
		database.DBClient.Model(&model.User{}).Where("email = ?", strings.ToLower(bodyData.Email)).Or("username = ?", bodyData.Username).Find(&userExist)

		if !userExist {
			helpers.RespondWithError(w, http.StatusBadRequest, "email or username already exist")
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
			helpers.RespondWithJSON(w, http.StatusCreated, userData)
		}
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// check content type
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		// response error
		helpers.RespondWithError(w, http.StatusUnsupportedMediaType, "Content Type is not application/json")
		return
	}
	// login either username or email and password
	var loginBody model.LoginBody
	// decode body data
	json.NewDecoder(r.Body).Decode(&loginBody)
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
		helpers.RespondWithError(w, http.StatusBadRequest, customValidationErr)
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
				database.DBClient.Model(&userdata).Updates(model.User{Token: authToken, RefreshToken: refreshToken}).Where("id = ?", userdata.ID)
				userdata.Token = authToken
				userdata.RefreshToken = refreshToken
				helpers.RespondWithJSON(w, http.StatusOK, userdata)
			} else {
				// password don't match
				helpers.RespondWithError(w, http.StatusBadRequest, "incorrect password")
			}
		} else {
			// user not found
			helpers.RespondWithError(w, http.StatusForbidden, "user not found")
		}
	}
}
