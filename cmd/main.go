package main

import (
	"github.com/TulioGuaraldoB/pizoo/internal/config/env"
	"github.com/TulioGuaraldoB/pizoo/internal/domain/usecase"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/handler"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/postgres"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/postgres/repository"
	"github.com/TulioGuaraldoB/pizoo/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.LogInfo("RUNNING")

	// Environment Variable
	env.GetEnvironmentVariable()

	// Driven
	db := postgres.OpenConnection()

	// Repositories
	animalRepository := repository.NewAnimalRepository(db)

	// Usecases
	animalUsecase := usecase.NewAnimalUsecase(animalRepository)

	// Handlers
	animalHandler := handler.NewAnimalHandler(animalUsecase)

	app := fiber.New()
	app.Get("/api/animal/", animalHandler.GetAllAnimals)
	app.Get("/api/animal/:id", animalHandler.GetAnimalById)
	app.Post("/api/animal/", animalHandler.CreateAnimal)
	app.Patch("/api/animal/:id", animalHandler.DeleteAnimalById)

	app.Listen(":" + env.Env.Port)
}
