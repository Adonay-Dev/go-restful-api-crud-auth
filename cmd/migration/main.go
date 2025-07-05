package main

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/logger"
	"os"
)

func main() {
	file, err := os.Create("../../logs/eventManagment.log")
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

	log.Infof("Loaded config: host=%s, port=%s, name=%s", cfg.DBHost, cfg.DBPort, cfg.DBName)

}
