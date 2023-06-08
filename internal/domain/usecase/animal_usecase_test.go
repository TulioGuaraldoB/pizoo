package usecase_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/TulioGuaraldoB/pizoo/internal/domain/entity"
	"github.com/TulioGuaraldoB/pizoo/internal/domain/usecase"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/postgres/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAnimalsUsecase(t *testing.T) {
	mockAnimal := new(entity.Animal)
	mockAnimal.ID = 1
	mockAnimal.Name = "zeze"
	mockAnimal.Breed = "asdfasdfasdf"
	mockAnimal.Gender = "sdafasdfasasdfdsa"
	mockAnimal.Age = 3
	mockAnimal.Size = "asdfasdfasdf"
	mockAnimal.City = "asdfasdf"
	mockAnimal.State = "qwetwqert"
	mockAnimal.Dewormed = "wqer"
	mockAnimal.Castrated = "qwe"
	mockAnimal.Vaccinated = "nbv"
	mockAnimal.SpecialCare = "sadf"
	mockAnimal.Picture = "sda"
	mockAnimal.CreatedAt = time.Now()
	mockAnimal.UpdatedAt = time.Now()
	mockAnimal.DeletedAt = sql.NullTime{}

	mockAnimals := make([]entity.Animal, 0)
	mockAnimals = append(mockAnimals, *mockAnimal)

	tests := []struct {
		description   string
		setMocks      func(mar *mock.MockIAnimalRepository)
		isErrExpected bool
	}{
		{
			description: "should return no error",
			setMocks: func(mar *mock.MockIAnimalRepository) {
				mar.EXPECT().GetAllAnimals().Return(mockAnimals, nil)
			},
		},
	}

	for _, tc := range tests { // testCase
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mar := mock.NewMockIAnimalRepository(controller)
			tc.setMocks(mar)

			// Act
			animalUsecase := usecase.NewAnimalUsecase(mar)
			_, err := animalUsecase.GetAllAnimals()
			// Assert
			if !tc.isErrExpected {
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}
}
