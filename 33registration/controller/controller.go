package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kiransharma121/registration/database"
	"github.com/kiransharma121/registration/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the serveHome"))
}
func hashingPassword(password string) (string, error) {
	Hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in hashedpassword")
	}
	return string(Hashedpassword), nil
}
func signUp(userInfo model.Registration) {
	hasedString, err := hashingPassword(userInfo.Password)
	if err != nil {
		fmt.Println("error in hasedString")
	}
	userInfo.Password = hasedString
	inserted, err := database.Collection.InsertOne(context.Background(), userInfo)
	if err != nil {
		fmt.Println("Error in inserting the user")
	}
	fmt.Println("Inserted data with the id:", inserted.InsertedID)
}
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var user model.Registration
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter the data")
	}
	json.NewDecoder(r.Body).Decode(&user)
	if r.Body == nil {
		fmt.Println("error")
	}
	signUp(user)
	json.NewEncoder(w).Encode(user)
}
func compareHashedPassword(hased, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hased), []byte(password))
	return err == nil
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method == http.MethodPost {
		var login loginRequest
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			fmt.Println("Invalid request")
		}
		var user model.Registration
		err = database.Collection.FindOne(context.Background(), bson.M{"email": login.Email}).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid value", http.StatusBadGateway)
		}
		if compareHashedPassword(login.Password, user.Password) {
			w.Write([]byte("user found"))
		} else {
			w.Write([]byte("user not found"))
		}

	}
}
