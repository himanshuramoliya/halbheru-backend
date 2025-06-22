package main

import (
	"log"
	"os"

	"halbheru-backend/database"
	"halbheru-backend/handlers"
	"halbheru-backend/middleware"

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

	// Initialize database connection
	// database.Connect()
	// database.Migrate()

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

	// Routes
	setupRoutes(app)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func setupRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Halbheru Backend is running",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Authentication routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes - require authentication
	protected := api.Use(middleware.AuthRequired())

	// Ride routes
	rides := protected.Group("/rides")
	rides.Get("/", getRidesHandler)
	rides.Post("/", createRideHandler)
	rides.Get("/:id", getRideHandler)
	rides.Put("/:id", updateRideHandler)
	rides.Delete("/:id", deleteRideHandler)
	rides.Post("/:id/join", joinRideHandler)

	// User routes
	users := protected.Group("/users")
	users.Get("/profile", getUserProfileHandler)
	users.Put("/profile", updateUserProfileHandler)
}

// Placeholder handlers - to be implemented
func getRidesHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get rides endpoint"})
}

func createRideHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create ride endpoint"})
}

func getRideHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get ride endpoint"})
}

func updateRideHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update ride endpoint"})
}

func deleteRideHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete ride endpoint"})
}

func joinRideHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Join ride endpoint"})
}

func getUserProfileHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get user profile endpoint"})
}

func updateUserProfileHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update user profile endpoint"})
}
