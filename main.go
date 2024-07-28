package main

import (
	"context"
	"log"

	"net/http"
	"os"

	// "strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDo struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var Collection *mongo.Collection

func main() {

	log.Printf("Hello There")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error while loadin .env file:", err)
	}

	var PORT string = os.Getenv("PORT")
	var MongoDB_URI string = os.Getenv("MongoDB_URI")

	clientOptions := options.Client().ApplyURI(MongoDB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection succesfull")

	Collection = client.Database("todoList_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", func(c *fiber.Ctx) error {

		var todo_list []ToDo

		cursor, err := Collection.Find(context.Background(), bson.M{})

		if err != nil {
			log.Fatal(err)
		}

		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var todo ToDo
			if err := cursor.Decode(&todo); err != nil {
				return err
			}
			todo_list = append(todo_list, todo)
		}

		return c.Status(http.StatusOK).JSON(todo_list)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {

		todo := new(ToDo)

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "todo body can not be empty"})
		}

		insertResult, err := Collection.InsertOne(context.Background(), todo)

		if err != nil {
			return err
		}

		todo.Id = insertResult.InsertedID.(primitive.ObjectID)

		return c.Status(http.StatusCreated).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return err
		}

		filter := bson.M{"_id": objectID}
		update := bson.M{"$set": bson.M{"completed": true}}

		_, err = Collection.UpdateOne(context.Background(), filter, update)

		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "task ended succesfully"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return err
		}

		filter := bson.M{"_id": objectID}

		deleteResult, err := Collection.DeleteOne(context.Background(), filter)

		if err != nil {
			return err
		}

		if deleteResult.DeletedCount == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "task was not found"})
		}

		return c.Status(http.StatusOK).JSON(deleteResult)
	})

	log.Fatal(app.Listen(":" + PORT))
}

// =====================

// go get go.mongodb.org/mongo-driver/mongo

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
//   )

//   func main() {
// 	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI("mongodb+srv://kshahkolaee:oSAA7qZDZg7xJnqj@cluster0.8vvtryx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 	  panic(err)
// 	}

// 	defer func() {
// 	  if err = client.Disconnect(context.TODO()); err != nil {
// 		panic(err)
// 	  }
// 	}()

// 	// Send a ping to confirm a successful connection
// 	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
//   }
