package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envKeys = []string{"DB_USER_NAME", "DB_PASSWORD", "DB_NAME", "PORT"}

func init() {
	err := godotenv.Load("../env/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for _, key := range envKeys {
		if _, ok := os.LookupEnv(key); !ok {
			fmt.Printf("%v env not loaded\n", key)
			os.Exit(1)
		}
	}
	fmt.Println("environment variables loaded successfully")
}

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	app.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
