package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Cosmos-Harry/blockchain-qa/indexer/internal/database"
	"github.com/Cosmos-Harry/blockchain-qa/indexer/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	ctx := context.Background()

	// Connect to database
	db, err := database.NewDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Connect to Redis (optional)
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Test Redis connection
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Redis not available: %v", err)
		redisClient = nil
	} else {
		log.Println("Connected to Redis")
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Blockchain QA API",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Initialize handlers
	pollHandler := handlers.NewPollHandler(db, redisClient)

	// Routes
	api := app.Group("/api")

	// Poll routes
	polls := api.Group("/polls")
	polls.Get("/", pollHandler.ListPolls)
	polls.Get("/:address", pollHandler.GetPoll)
	polls.Get("/:address/votes", pollHandler.GetPollVotes)
	polls.Get("/:address/results", pollHandler.GetPollResults)
	polls.Get("/:address/stats", pollHandler.GetVoteCount)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now().UTC().Format(time.RFC3339),
		})
	})

	// Start server in a goroutine
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		log.Printf("Starting API server on port %s\n", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Gracefully shutdown the server
	if err := app.Shutdown(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}

	log.Println("Server stopped")
}
