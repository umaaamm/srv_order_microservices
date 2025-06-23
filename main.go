package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"srv_order/main/api/router"
	contact "srv_order/main/pkg/order"
	pb "srv_order/main/proto/contact"
)

func main() {
	db, dbCancel, err := databaseConnection()
	if err != nil {
		log.Fatalf("Database Connection Error: %v", err)
	}
	defer dbCancel()
	fmt.Println("Database connection success!")

	conn, err := grpc.Dial("service-contact:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to user-service: %v", err)
	}
	defer conn.Close()

	userClient := pb.NewContactServiceClient(conn)

	contactCollection := db.Collection("order")
	contactRepo := contact.NewRepo(contactCollection)
	contactService := contact.NewService(contactRepo, userClient)

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Services Contact API is running")
	})

	api := app.Group("/api")
	router.ContactRouter(api, contactService)

	log.Println("Fiber API running on :8082")
	log.Fatal(app.Listen(":8082"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	uri := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("contact")
	return db, cancel, nil
}
