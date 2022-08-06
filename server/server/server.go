package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faysal146/golang-vue-http-crud-app/server/server/routes"
)

const _PORT = 8080

func StartServer() {

	routersHandler := routes.RoutesMux()

	host := fmt.Sprintf(":%v", _PORT)
	fmt.Println("Server Running on Port ", _PORT)
	err := http.ListenAndServe(host, routersHandler)
	if err != nil {
		log.Fatal(err)
	}
}
