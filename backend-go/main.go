package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type PlanResponse struct {
	Strategy string `json:"strategy"`
}

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	aiURL := os.Getenv("AI_SERVICE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// AI plan proxy
	app.Get("/plan", func(c *fiber.Ctx) error {
		resp, err := http.Get(fmt.Sprintf("%s/plan", aiURL))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to reach AI service", "details": err.Error()})
		}
		defer resp.Body.Close()

		var plan PlanResponse
		if err := json.NewDecoder(resp.Body).Decode(&plan); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "invalid response from AI"})
		}

		return c.JSON(plan)
	})

	log.Printf("Backend running on port %s", port)
	app.Listen(fmt.Sprintf(":%s", port))
}
