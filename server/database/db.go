package database

import (
	"fmt"
	"log"
	"os"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func InitializeDB(user, password, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Dhaka", user, password, dbname, 5432)

	var err error
	DBClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect with database", err)
		os.Exit(1)
	} else {
		fmt.Println("database connected...")
	}
	// Migrate database
	if userDBMigErr := model.MigrateUser(DBClient); userDBMigErr != nil {
		log.Fatal("user database Migrate fail ", err)
	} else {
		fmt.Println("user database Migrate successfully")
	}
	return DBClient
}
