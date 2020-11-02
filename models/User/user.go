package User

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	username string
	password string
	aircraft string
}
