package main

import (
	"flag"
	"fmt"
	"go-rest-api/internal/config"
	"go-rest-api/internal/database"
	"go-rest-api/internal/logger"
	"os"
	"path/filepath"

	"github.com/pressly/goose"
)

func main() {
	logPath := filepath.Join("..", "logs", "eventManagment.log")
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger.Init(file)

	log := logger.Get()

	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Errorf("Configuration loading error: %v", err)
	}

	db, err := database.DBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	migrate := flag.Bool("migrate", false, "Run database migrations and exit")
	dir := flag.String("dir", "../migrations", "Directory with the migration files")
	action := flag.String("action", "up", "migration action: up, down, status")
	flag.Parse()

	if *migrate {
		switch *action {
		case "up":
			if err := goose.Up(db, *dir); err != nil {
				log.Fatalf("goose up failed: %v", err)
			}
			fmt.Println("Migration up completed.")
		case "down":
			if err := goose.Down(db, *dir); err != nil {
				log.Fatalf("goose down failed: %v", err)
			}
			fmt.Println("Migration down completed.")
		case "status":
			if err := goose.Status(db, *dir); err != nil {
				log.Fatalf("goose status failed: %v", err)
			}
		default:
			log.Fatalf("Unknown action: %s", *action)
		}
	}
	startServer()
}

func startServer() {
	fmt.Print("Server Running...")
}
