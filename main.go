package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/Ledja22/hotel-reservation/api"
	"github.com/Ledja22/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "James",
		LastName:  "At the water cooler thot",
	}

	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("error when inserting a user")
	}

	var james types.User

	if err := coll.FindOne(ctx, bson.M{}).Decode(&james); err != nil {
		fmt.Println("error when finding a user")
	}

	fmt.Println(james)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New()
	appv1 := app.Group("/api/v1")
	appv1.Get("/user", api.HandleGetUser)
	appv1.Get("/users/", api.HandleGetUsers)
	app.Listen(*listenAddr)
}
