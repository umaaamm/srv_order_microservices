package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"srv_contact/main/api/router"
	"srv_contact/main/pkg/contact"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	contactCollection := db.Collection("contacts")
	contactRepo := contact.NewRepo(contactCollection)
	contactService := contact.NewService(contactRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Services Contact API is running"))
	})
	api := app.Group("/api")
	router.ContactRouter(api, contactService)
	defer cancel()
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://username:password@localhost:27017/contact").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("contact")
	return db, cancel, nil
}
