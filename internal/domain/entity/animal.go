package entity

import (
	"database/sql"
	"time"
)

type Animal struct {
	ID          uint
	Name        string
	Breed       string
	Gender      string
	Age         int64
	Size        string
	City        string
	State       string
	Dewormed    string
	Castrated   string
	Vaccinated  string
	SpecialCare string
	Picture     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type IAnimalRepository interface {
	GetAllAnimals() ([]Animal, error)
	GetAnimalById(animalId uint) (*Animal, error)
	CreateAnimal(animal *Animal) error
	DeleteAnimal(animalId uint) error
}
