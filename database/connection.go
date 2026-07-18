package database

import "C"
import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var DB *mongo.Database
var Collection *mongo.Collection
var Context *mongo.Database
var Collection1 *mongo.Collection

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoURI := "mongodb://localhost:27017"

	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err == nil {
		//panic(err)
		fmt.Println("Connection to DB", Client)
		DB = Client.Database("social") // Db name declared in the mongoDB we will change it later
	} else {
		fmt.Println("err  to connect to  MongoDB", err.Error())
		//		fmt.Println("Eror connecting to DB", err.Error())
	}
	//
	//Collection = DB.Collection("airplay")
	collection := Client.Database("test").Collection("social")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Find failed:", err.Error())
		return
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println("Cursor error:", err.Error())
		return
	}

	for _, doc := range results {
		fmt.Println(doc)
	}

}
