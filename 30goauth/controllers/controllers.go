package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kiransharma121/goauth/database"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the server home"))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	newUser := bson.D{
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: hashedPassword},
	}
	fmt.Println(user.Username)
	fmt.Println(user.Email)
	fmt.Println(hashedPassword)

	_, err = database.Collection.InsertOne(context.Background(), newUser)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
	defer r.Body.Close()
}
