package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/Ledja22/hotel-reservation/api"
	"github.com/Ledja22/hotel-reservation/api/middleware"
	"github.com/Ledja22/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		fmt.Println(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New(config)
	appv1 := app.Group("/api/v1", middleware.JWTAuthentication)
	auth := app.Group("/api")
	hotelStore := db.NewMongoHotelStore(client, dbname)
	roomStore := db.NewMongoRoomStore(client, dbname, hotelStore)
	userStore := db.NewMongoUserStore(client, db.DBNAME)
	store := &db.Store{
		User:  userStore,
		Hotel: hotelStore,
		Room:  roomStore,
	}

	userHandler := api.NewUserHandler(userStore)
	hotelHandler := api.NewHotelHandler(store)
	authHandler := api.NewAuthHandler(userStore)

	auth.Post("/auth", authHandler.HandleAuthenticate)
	appv1.Put("/user/:id", userHandler.HandlePutUser)
	appv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	appv1.Post("/user", userHandler.HandlePostUser)
	appv1.Get("/user", userHandler.HandleGetUsers)
	appv1.Get("/user/:id", userHandler.HandleGetUser)

	appv1.Get("/hotel", hotelHandler.HandleGetHotels)
	appv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	appv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	app.Listen(*listenAddr)
}
