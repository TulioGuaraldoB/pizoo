package handler

import (
	"github.com/TulioGuaraldoB/pizoo/internal/domain/usecase"
	"github.com/kataras/iris/v12"
)

type IAnimalHandler interface {
	GetAllAnimals(c iris.Context)
	GetAnimalById(c iris.Context)
	CreateAnimal(c iris.Context)
	DeleteAnimalById(c iris.Context)
}

func NewAnimalHandler(animalUsecase usecase.IAnimalUsecase) IAnimalHandler {
	return &animalHandler{animalUsecase: animalUsecase}
}
