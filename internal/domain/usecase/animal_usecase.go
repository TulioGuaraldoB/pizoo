package usecase

import "github.com/TulioGuaraldoB/pizoo/internal/domain/entity"

type animalUsecase struct {
	animalRepository entity.IAnimalRepository
}

func (u *animalUsecase) GetAllAnimals() ([]entity.Animal, error) {
	return u.animalRepository.GetAllAnimals()
}

func (u *animalUsecase) GetAnimalById(animalId uint) (*entity.Animal, error) {
	return u.animalRepository.GetAnimalById(animalId)
}

func (u *animalUsecase) CreateAnimal(animal *entity.Animal) error {
	return u.animalRepository.CreateAnimal(animal)
}

func (u *animalUsecase) DeleteAnimal(animalId uint) error {
	return u.animalRepository.DeleteAnimal(animalId)
}
