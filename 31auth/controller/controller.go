package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kiransharma121/auth/database"
	"github.com/kiransharma121/auth/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the serve home"))
}
func hashingPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in hasing the password", err)
	}
	return string(hashed), nil
}
func insertOneUser(userinfo model.Authentication) {
	hashed, err := hashingPassword(userinfo.Password)
	if err != nil {
		panic(err)
	}
	userinfo.Password = hashed
	inserted, err := database.Collection.InsertOne(context.Background(), userinfo)
	if err != nil {
		fmt.Println("Error in inserting the data in database", err)
	}
	fmt.Println("Data is inserted in the database with id:", inserted.InsertedID)
}
func InsertOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	var userdata model.Authentication
	_ = json.NewDecoder(r.Body).Decode(&userdata)
	insertOneUser(userdata)
	json.NewEncoder(w).Encode(userdata)

}
func getAlluser() []primitive.M {
	curr, err := database.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println("Error in finding the data", err)
	}
	defer curr.Close(context.Background())
	var users []primitive.M
	for curr.Next(context.Background()) {
		var user bson.M
		err := curr.Decode(&user)
		if err != nil {
			fmt.Println("Error in decoding the user from database", err)
		}
		users = append(users, user)
	}
	return users
}
func Getallusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	user := getAlluser()
	json.NewEncoder(w).Encode(user)
}
