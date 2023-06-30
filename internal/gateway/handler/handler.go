package handler

import (
	"github.com/TulioGuaraldoB/pizoo/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

type IAnimalHandler interface {
	GetAllAnimals(ctx *fiber.Ctx) error
	GetAnimalById(ctx *fiber.Ctx) error
	CreateAnimal(ctx *fiber.Ctx) error
	DeleteAnimalById(ctx *fiber.Ctx) error
}

func NewAnimalHandler(animalUsecase usecase.IAnimalUsecase) IAnimalHandler {
	return &animalHandler{animalUsecase: animalUsecase}
}
