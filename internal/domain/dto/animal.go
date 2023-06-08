package dto

import (
	"database/sql"
	"time"

	"github.com/TulioGuaraldoB/pizoo/internal/domain/entity"
)

type AnimalResponse struct {
	ID          uint   `json:""`
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Gender      string `json:"gender"`
	Age         int64  `json:"age"`
	Size        string `json:"size"`
	City        string `json:"city"`
	State       string `json:"state"`
	Dewormed    string `json:"dewormed"`
	Castrated   string `json:"castrated"`
	Vaccinated  string `json:"vaccinated"`
	SpecialCare string `json:"special_care"`
	Picture     string `json:"picture"`
}

type AnimalRequest struct {
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Gender      string `json:"gender"`
	Age         int64  `json:"age"`
	Size        string `json:"size"`
	City        string `json:"city"`
	State       string `json:"state"`
	Dewormed    string `json:"dewormed"`
	Castrated   string `json:"castrated"`
	Vaccinated  string `json:"vaccinated"`
	SpecialCare string `json:"special_care"`
	Picture     string `json:"picture"`
}

func MapAnimalResponse(animal *entity.Animal) *AnimalResponse {
	return &AnimalResponse{
		ID:          animal.ID,
		Name:        animal.Name,
		Breed:       animal.Breed,
		Gender:      animal.Gender,
		Age:         animal.Age,
		Size:        animal.Size,
		City:        animal.City,
		State:       animal.State,
		Dewormed:    animal.Dewormed,
		Castrated:   animal.Castrated,
		Vaccinated:  animal.Vaccinated,
		SpecialCare: animal.SpecialCare,
		Picture:     animal.Picture,
	}
}

func MapAnimalRequest(animalRequest *AnimalRequest) *entity.Animal {
	return &entity.Animal{
		Name:        animalRequest.Name,
		Breed:       animalRequest.Breed,
		Gender:      animalRequest.Gender,
		Age:         animalRequest.Age,
		Size:        animalRequest.Size,
		City:        animalRequest.City,
		State:       animalRequest.State,
		Dewormed:    animalRequest.Dewormed,
		Castrated:   animalRequest.Castrated,
		Vaccinated:  animalRequest.Vaccinated,
		SpecialCare: animalRequest.SpecialCare,
		Picture:     animalRequest.Picture,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   sql.NullTime{},
	}
}
