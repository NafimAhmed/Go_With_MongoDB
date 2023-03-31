package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"golang.org/x/vuln/client"
)

const URL = "mongodb://localhost:27017"

// type Books struct {
// 	isbn   string
// 	title  string
// 	author string
// }

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))

	if err != nil {
		panic(err)
	}

	defer func() {

		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// books := Books{
	// 	isbn:   "123456",
	// 	title:  "dkdhkhd",
	// 	author: "ifjfjhf",
	// }

	collections := client.Database("First_Database").Collection("First_Collection")

	data := bson.D{
		{Key: "isbn", Value: "123456"},
		{Key: "title", Value: "dkdhkhd"},
		{Key: "author", Value: "ifjfjhf"},
	}

	resultset, err := collections.InsertOne(ctx, data)

	fmt.Println("Insert ID : %v", resultset.InsertedID)

}
