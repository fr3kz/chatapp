package utils

import (
	user_model "../models/User"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

var SecretKey = "value"

func ConnectTobDb() *gorm.DB {
	var db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=chat password=BIko0230 sslmode=disable")
	if err != nil {
		fmt.Print(err.Error())
	}
	db.AutoMigrate(&user_model.User{})
	return db
}

func RespondWithJson(w http.ResponseWriter, code int, message interface{}) {

	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(message)
	w.Write(payload)
}
