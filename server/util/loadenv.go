package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var EnvKeys = []string{"DB_USER_NAME", "DB_PASSWORD", "DB_NAME", "PORT", "JWT_SECRETKEY"}

func LoadEnv(envKeys []string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("unable to get the current filename")
		os.Exit(1)
	}
	dirname := filepath.Dir(filename)
	var err error
	if os.Getenv("GO_ENV") == "test" {
		err = godotenv.Load(
			filepath.Join(dirname, "../", "env/test.env"),
		)
	} else {
		err = godotenv.Load(filepath.Join(dirname, "../env/config.env"))
	}
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	for _, key := range envKeys {
		if _, ok := os.LookupEnv(key); !ok {
			fmt.Printf("%v env not loaded\n", key)
			os.Exit(1)
		}
	}
	fmt.Println("environment variables loaded successfully")
}
