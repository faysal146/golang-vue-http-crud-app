package routes

import (
	"net/http"

	"github.com/faysal146/golang-vue-http-crud-app/server/pkg/controllers"
	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router) {
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/auth/login", controllers.LoginUser).Methods(http.MethodPost)
	s.HandleFunc("/auth/register", controllers.RegisterUser).Methods(http.MethodPost)
}
