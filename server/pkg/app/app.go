package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/routes"
	"github.com/gorilla/mux"
)

func Run(addr string) {
	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	fmt.Println("server running on port ", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Println("server could not running...")
		log.Fatal(err)
	}
}
