package main

import (
	"aroftu/darsuka/ent"
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	client, dbErr := ent.Open("sqlite3", "darsuka.db?_fk=1")

	if dbErr != nil {
		log.Fatalf("failed opening sqlite database: %v", dbErr)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		ctx := context.Background()
		users, err := client.User.Query().All(ctx)
		if err != nil {
			return c.JSON(fiber.Map{"data": "error"})
		}
		return c.JSON(fiber.Map{"data": users})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		user := ent.User{}
		if err := c.BodyParser(&user); err != nil {
			log.Fatalf("failed to parse body: %v", err)
		}
		client.User.Create().
			SetEmail(user.Email).
			SetName(user.Name).
			SetPassword(user.Password).
			Save(context.Background())

		return c.JSON(fiber.Map{"msg": "success add"})
	})

	app.Listen(":3000")
}
