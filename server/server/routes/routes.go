package routes

import (
	"net/http"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/controller"
	"github.com/gorilla/mux"
)

func RoutesMux() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.HelloWorld).Methods(http.MethodGet)

	return r
}
