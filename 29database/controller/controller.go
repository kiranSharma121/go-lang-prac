package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiransharma121/mangodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "Employee"
const colName = "presentList"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongodb connected sucessfully...")
	collection = client.Database(dbName).Collection(colName)

}
func insertFullName(fullname model.Employe) {
	inserted, err := collection.InsertOne(context.Background(), fullname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The inserted firstname in database with id:", inserted.InsertedID)
}
func updateOneName(nameID string) {
	id, err := primitive.ObjectIDFromHex(nameID)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"present": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Modified counter:", result.MatchedCount)

}
func deleteOneName(nameID string) {
	id, err := primitive.ObjectIDFromHex(nameID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Movie has been deleted with id :", result)
}
func deleteAllName() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("All name are deleted:", result.DeletedCount)
}
func getAllName() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	var Names []primitive.M
	for cur.Next(context.Background()) {
		var Name bson.M
		err := cur.Decode(&Name)
		if err != nil {
			panic(err)
		}
		Names = append(Names, Name)
	}
	defer cur.Close(context.Background())
	return Names

}
func GetAllName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	getallname := getAllName()
	json.NewEncoder(w).Encode(getallname)
}
func CreateName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var name model.Employe
	_ = json.NewDecoder(r.Body).Decode(&name)
	insertFullName(name)
	json.NewEncoder(w).Encode(name)
}
func UpDateOneName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneName(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	json.NewEncoder(w).Encode(deleteAllName)
}
func DeleteOneName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneName(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
