package server

import (
	authcontroller "../controllers/Auth"
	a "../models/User"
	"../utils"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/login", authcontroller.Login)
	r.HandleFunc("/register", authcontroller.Register)
	http.ListenAndServe(":8080", r)
}

func StartServerDebug() {
	//migrates
	db := utils.ConnectTobDb()
	db.AutoMigrate(&a.User{})
	db.Close()
	StartServer()
}
