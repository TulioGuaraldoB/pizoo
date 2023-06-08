package usecase

import "github.com/TulioGuaraldoB/pizoo/internal/domain/entity"

type IAnimalUsecase interface {
	GetAllAnimals() ([]entity.Animal, error)
	GetAnimalById(animalId uint) (*entity.Animal, error)
	CreateAnimal(animal *entity.Animal) error
	DeleteAnimal(animalId uint) error
}

func NewAnimalUsecase(animalRepository entity.IAnimalRepository) IAnimalUsecase {
	return &animalUsecase{animalRepository: animalRepository}
}
