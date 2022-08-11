package main

import (
	"fmt"
	"os"

	"github.com/faysal146/golang-vue-http-crud-app/server/database"
	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/app"
	"github.com/faysal146/golang-vue-http-crud-app/server/util"
)

func init() {
	// load environment variables
	util.LoadEnv(util.EnvKeys)
}

func main() {
	// connect database
	database.InitializeDB(
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// start server
	app.Run(fmt.Sprintf("127.0.0.1:%v", os.Getenv("PORT")))
}
