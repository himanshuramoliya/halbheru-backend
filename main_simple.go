package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Halbheru Backend API v1.0.0",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "*",
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Halbheru Backend is running",
			"version": "1.0.0",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Test endpoints
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API is working!",
			"endpoints": []string{
				"GET /health",
				"GET /api/v1/test",
				"POST /api/v1/auth/register",
				"POST /api/v1/auth/login",
			},
		})
	})

	// Authentication routes (mock)
	auth := api.Group("/auth")
	auth.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Registration endpoint ready",
			"note":    "Database connection needed for full functionality",
		})
	})
	
	auth.Post("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Login endpoint ready",
			"note":    "Database connection needed for full functionality",
		})
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📋 Health check: http://localhost:%s/health", port)
	log.Printf("🧪 Test endpoint: http://localhost:%s/api/v1/test", port)
	log.Fatal(app.Listen(":" + port))
}
