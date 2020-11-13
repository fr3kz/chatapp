package Auth

import (
	"../../models/User"
	"../../utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var users []User.User

func Login(w http.ResponseWriter, r *http.Request) {
	//perform auth check
	fmt.Print(r.FormValue("password"))
	db := utils.ConnectTobDb()
	db.Where(map[string]interface{}{"username": r.FormValue("username"), "password": r.FormValue("password")}).Find(&users)
	db.Close()

	if len(users) > 0 {
		// auth confirmed
		Authenticate(w, r.FormValue("username"))
	} else {
		//not authed
		utils.RespondWithJson(w, http.StatusBadRequest, "error: bad cred")
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	user := User.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	db := utils.ConnectTobDb()
	db.Create(&user)
	db.Close()

	Authenticate(w, user.Username)
}

func Token_Check(r *http.Request) jwt.MapClaims {
	tokenString := r.Header.Get("key")
	result, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SecretKey), nil
	})
	c := result.Claims.(jwt.MapClaims)
	//c is a map[string]string [username]username_value
	return c
}

func Authenticate(w http.ResponseWriter, username string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	token_string, err := token.SignedString([]byte(utils.SecretKey))

	if err != nil {
		utils.RespondWithJson(w, http.StatusBadRequest, "error: "+err.Error())
	}
	utils.RespondWithJson(w, http.StatusAccepted, "token: "+token_string)
}
