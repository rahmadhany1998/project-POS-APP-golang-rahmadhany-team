package main

import (
	"log"
	"project-POS-APP-golang-be-team/cmd"
	"project-POS-APP-golang-be-team/internal/data"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/wire"
	"project-POS-APP-golang-be-team/pkg/database"
	"project-POS-APP-golang-be-team/pkg/middleware"
	"project-POS-APP-golang-be-team/pkg/utils"

	"go.uber.org/zap"
)

func main() {

	// read config
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	logger, err := utils.InitLogger(config.PathLogger, config)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}

	//Init db
	db, err := database.InitDB(config)
	if err != nil {
		logger.Fatal("can't connect to database ", zap.Error(err))
	}

	// migration
	if err := data.AutoMigrate(db); err != nil {
		logger.Fatal("failed to run migrations", zap.Error(err))
	}

	// seeder
	if err := data.SeedAll(db); err != nil {
		logger.Fatal("failed to seed initial data", zap.Error(err))
	}

	repo := repository.NewRepository(db, logger)
	mLogger := middleware.NewLoggerMiddleware(logger)
	mAuth := middleware.NewAuthMiddleware(repo, logger)
	router := wire.Wiring(repo, mLogger, mAuth, logger, config)

	cmd.ApiServer(config, logger, router)
}
