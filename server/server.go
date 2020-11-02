package server

import (
	"fmt"
	"github.com/gorilla/mux"
	authcontroller "../controllers/Auth"
	"net/http"
	"../utils"
	user_model "../models/User"
)

func StartServer(){
	r := mux.NewRouter()
	r.HandleFunc("/login",authcontroller.Login)
	http.ListenAndServe(":8080",r)
}

func StartServerDebug(){
	//migrates
	db := utils.ConnectTobDb()
	fmt.Print(db)
	db.AutoMigrate(&user_model.UserModel{})
	StartServer()
}