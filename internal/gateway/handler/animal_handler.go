package handler

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/pizoo/internal/domain/dto"
	"github.com/TulioGuaraldoB/pizoo/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

type animalHandler struct {
	animalUsecase usecase.IAnimalUsecase
}

func (h *animalHandler) GetAllAnimals(ctx *fiber.Ctx) error {
	animals, err := h.animalUsecase.GetAllAnimals()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(animals) == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "record not found"})
	}

	animalsResponse := make([]dto.AnimalResponse, 0)
	for _, animal := range animals {
		animalsResponse = append(animalsResponse, *dto.MapAnimalResponse(&animal))
	}

	return ctx.Status(http.StatusOK).JSON(animalsResponse)
}

func (h *animalHandler) GetAnimalById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	animal, err := h.animalUsecase.GetAnimalById(uint(id))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if animal.ID <= 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "record not found"})
	}

	return ctx.Status(http.StatusOK).JSON(dto.MapAnimalResponse(animal))
}

func (h *animalHandler) CreateAnimal(ctx *fiber.Ctx) error {
	animalRequest := new(dto.AnimalRequest)
	if err := ctx.BodyParser(animalRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.animalUsecase.CreateAnimal(dto.MapAnimalRequest(animalRequest)); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusCreated).
		JSON(fiber.Map{
			"success": "animal created!",
			"animal":  animalRequest,
		})
}

func (h *animalHandler) DeleteAnimalById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.animalUsecase.DeleteAnimal(uint(id)); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"success": "animal removed!"})
}
