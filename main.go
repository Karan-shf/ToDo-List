package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ToDo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()
	log.Printf("Hello There")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error while loadin .env file")
	}

	var PORT string = os.Getenv("PORT")

	todo_list := []ToDo{}

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.Status(http.StatusOK).JSON(fiber.Map{"message": "hello there my friend"})
		return c.Status(http.StatusOK).JSON(todo_list)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// Todo := ToDo{}
		var Todo = ToDo{}

		if err := c.BodyParser(&Todo); err != nil {
			return err
		}

		if Todo.Body == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Todo body is required"})
		}

		Todo.Id = len(todo_list) + 1
		todo_list = append(todo_list, Todo)

		return c.Status(http.StatusCreated).JSON(Todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {

		idstr := c.Params("id")

		id, _ := strconv.Atoi(idstr)

		for index, todo := range todo_list {
			if todo.Id == id {
				todo_list[index].Completed = true
				return c.Status(http.StatusOK).JSON(todo_list[index])
			}
		}

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "task not found"})

	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		idstr := c.Params("id")

		id, _ := strconv.Atoi(idstr)

		for index, todo := range todo_list {
			if todo.Id == id {
				todo_list = append(todo_list[:index], todo_list[index+1:]...)
				return c.Status(http.StatusOK).JSON(fiber.Map{"message": "taks deleted succesfully"})
			}
		}

		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
