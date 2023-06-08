package handler

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/pizoo/internal/domain/dto"
	"github.com/TulioGuaraldoB/pizoo/internal/domain/usecase"
	"github.com/kataras/iris/v12"
)

type animalHandler struct {
	animalUsecase usecase.IAnimalUsecase
}

func (h *animalHandler) GetAllAnimals(c iris.Context) {
	animals, err := h.animalUsecase.GetAllAnimals()
	if err != nil {
		c.StatusCode(http.StatusInternalServerError)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	if len(animals) == 0 {
		c.StatusCode(http.StatusNotFound)
		c.JSON(iris.Map{"error": "record not found"})
		return
	}

	animalsResponse := make([]dto.AnimalResponse, 0)
	for _, animal := range animals {
		animalsResponse = append(animalsResponse, *dto.MapAnimalResponse(&animal))
	}
	c.StatusCode(http.StatusOK)
	c.JSON(animalsResponse)
}

func (h *animalHandler) GetAnimalById(c iris.Context) {
	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)
	if err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	animal, err := h.animalUsecase.GetAnimalById(uint(id))
	if err != nil {
		c.StatusCode(http.StatusInternalServerError)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	if animal.ID <= 0 {
		c.StatusCode(http.StatusNotFound)
		c.JSON(iris.Map{"error": "record not found"})
		return
	}

	c.StatusCode(http.StatusOK)
	c.JSON(dto.MapAnimalResponse(animal))
}

func (h *animalHandler) CreateAnimal(c iris.Context) {
	animalRequest := new(dto.AnimalRequest)
	if err := c.ReadJSON(animalRequest); err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	if err := h.animalUsecase.CreateAnimal(dto.MapAnimalRequest(animalRequest)); err != nil {
		c.StatusCode(http.StatusInternalServerError)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	c.StatusCode(http.StatusCreated)
	c.JSON(iris.Map{
		"success": "animal created!",
		"animal":  animalRequest,
	})
}

func (h *animalHandler) DeleteAnimalById(c iris.Context) {
	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)
	if err != nil {
		c.StatusCode(http.StatusBadRequest)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	if err := h.animalUsecase.DeleteAnimal(uint(id)); err != nil {
		c.StatusCode(http.StatusInternalServerError)
		c.JSON(iris.Map{"error": err.Error()})
		return
	}

	c.StatusCode(http.StatusOK)
	c.JSON(iris.Map{"success": "animal removed!"})
}
