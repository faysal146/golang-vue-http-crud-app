package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     any
}

func (a *App) Run(addr string) {
	fmt.Println(addr)
}

func (a *App) Initialize(user, password, dbname string) {
	fmt.Println(user, password, dbname)
}
