package main

import (
	"fmt"
	"os"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/app"
	"github.com/faysal146/golang-vue-http-crud-app/server/util"
)

func init() {
	util.LoadEnv(util.EnvKeys)
}

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	a.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
