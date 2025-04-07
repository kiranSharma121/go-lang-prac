package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiransharma121/login/database"
	"github.com/kiransharma121/login/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the serve Home"))
}
func insertOneUser(userinfo model.Auth) {
	hasedpassword, err := Hashedpassword(userinfo.Password)
	if err != nil {
		fmt.Println("Error in hasedpassword")
	}
	userinfo.Password = hasedpassword
	inserted, err := database.Collection.InsertOne(context.Background(), userinfo)
	if err != nil {
		fmt.Println("Error in inserting data in the database", err)
	}
	fmt.Println("Data has been inserted in the database with id", inserted.InsertedID)
}
func Hashedpassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in hassing the password")
	}
	return string(hashed), nil
}

func Insertoneuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var userdata model.Auth
	json.NewDecoder(r.Body).Decode(&userdata)
	insertOneUser(userdata)
	json.NewEncoder(w).Encode(userdata)
}
func getAllUser() []primitive.M {
	curr, err := database.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println("Error in finding the datafrom the database")
	}
	defer curr.Close(context.Background())
	var users []primitive.M
	for curr.Next(context.Background()) {
		var user bson.M
		err := curr.Decode(&user)
		if err != nil {
			fmt.Println("Error in decoding the user", err)
		}
		users = append(users, user)
	}
	return users

}
func Getalluser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	user := getAllUser()
	json.NewEncoder(w).Encode(user)

}
func updateOneUser(userid string) {
	id, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		fmt.Println("Error in getting the id", err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"userexist": true}}
	result, err := database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err, http.StatusBadRequest)
	}
	fmt.Println("The data has been updated", result.MatchedCount)

}
func Updateone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	updateOneUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func deleteOneUser(userid string) {
	id, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		fmt.Println("Error in getting the userid...x2")
	}
	filter := bson.M{"_id": id}
	result, err := database.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted the data with id", result.DeletedCount)
}
func Deleteoneuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	deleteOneUser(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var user model.Auth
		err := database.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		if checkPasswordHash(password, user.Password) {
			fmt.Fprintf(w, "login sucessful")
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)

		}
	}
}
