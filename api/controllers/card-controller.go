package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"

	"github.com/BackendApiCardExam/api/config"
	"github.com/BackendApiCardExam/api/models"
	"github.com/BackendApiCardExam/api/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//var NewCard models.Card

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.Card{}, &models.Comment{}, &models.QuestionAnswer{}, &models.User{})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetCardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["CardId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	var getCard models.Card
	db.Preload("QuestionsAnswers").Preload("Comments").Find(&getCard, "ID = ?", ID)
	res, _ := json.Marshal(getCard)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	CreateCard := &models.Card{}
	utils.ParseBody(r, CreateCard)
	db.Create(CreateCard)
	res, _ := json.Marshal(CreateCard)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCardAuthUserOnly(w http.ResponseWriter, r *http.Request) {
	secretkey := "my-secret-key"
	var mySigningKey = []byte(secretkey)
	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		//check if token is parsed correctly
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token wasn't signed with HMAC method")
		}
		return mySigningKey, nil
	})
	if err != nil {
		err := "error while parsing token"
		json.NewEncoder(w).Encode(err)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role"] == "user" {
			CreateCard := &models.Card{}
			utils.ParseBody(r, CreateCard)
			db.Create(CreateCard)
			res, _ := json.Marshal(CreateCard)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			res, _ := json.Marshal("not user")
			w.Write(res)
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var parsedUser models.User
	utils.ParseBody(r, &parsedUser)
	var newUser models.User
	db.Where("email = ?", parsedUser.Email).First(&newUser)

	if newUser.Email != "" {
		err := errors.New("email already in use")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	parsedUser.Password = utils.GeneratePassword(parsedUser.Password)

	db.Create(&parsedUser)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(parsedUser)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var authDetails models.Authentication
	utils.ParseBody(r, &authDetails)
	var loggedUser models.User
	db.Where("email = ?", authDetails.Email).Or("username = ?", authDetails.Username).First(&loggedUser)

	if loggedUser.Email == "" && loggedUser.Username == "" {
		err := errors.New("no such email or username")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	checkPass := utils.CheckPasswordHash(authDetails.Password, loggedUser.Password)

	if !checkPass {
		err := errors.New("username or password incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validToken := utils.GenerateJWT(loggedUser.Email, loggedUser.Role)
	var currentToken models.Token
	currentToken.TokenString = validToken
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currentToken)
}
