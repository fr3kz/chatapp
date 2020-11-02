package utils

import (
	user_model "../models/User"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectTobDb() *gorm.DB {
	var db,err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=chat password=BIko0230 sslmode=disable")
	fmt.Print(db,err)
	if err != nil{
		fmt.Print(err.Error())
	}
	db.AutoMigrate(&user_model.UserModel{})
	return db
}

func RespondWithJson(){}


