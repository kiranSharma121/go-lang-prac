package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

const connectionString = "mongodb://localhost:27017"
const dbName = "Authentication"
const colName = "Userdata"

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongodb connected sucessfully...")
	Collection = client.Database(dbName).Collection(colName)
}
