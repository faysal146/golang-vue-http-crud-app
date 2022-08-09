package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Run(addr string) {
	fmt.Println("server running on port ", addr)
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		fmt.Println("server could not running...")
		log.Fatal(err)
	}
}

func (a *App) Initialize(user, password, dbname string) {
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Dhaka", user, password, dbname, 5432)

	var err error
	a.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect with database", err)
		os.Exit(1)
	} else {
		fmt.Println("database connected...")
	}
	// Migrate database
	if userDBMigErr := model.MigrateUser(a.DB); userDBMigErr != nil {
		log.Fatal("user database Migrate fail ", err)
	} else {
		fmt.Println("user database Migrate successfully")
	}

	a.Router = mux.NewRouter()
}
