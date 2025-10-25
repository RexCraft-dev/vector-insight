package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Handicap  float32   `json:"handicap"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	app := fiber.New()

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Get all users
	app.Get("/users", func(c *fiber.Ctx) error {
		rows, err := dbpool.Query(ctx, "SELECT id, name, email, handicap, created_at FROM users ORDER BY id")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Handicap, &u.CreatedAt); err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			users = append(users, u)
		}

		return c.JSON(users)
	})

	// Create a user
	app.Post("/users", func(c *fiber.Ctx) error {
		var u User
		if err := c.BodyParser(&u); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
		}

		query := `INSERT INTO users (name, email, handicap) VALUES ($1, $2, $3) RETURNING id, created_at`
		err := dbpool.QueryRow(ctx, query, u.Name, u.Email, u.Handicap).Scan(&u.ID, &u.CreatedAt)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(u)
	})

	log.Println("Backend connected to DB and running on :8080")
	app.Listen(":8080")
}
