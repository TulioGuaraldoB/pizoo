package repository_test

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/TulioGuaraldoB/pizoo/internal/domain/entity"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/postgres/repository"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func mockDatabase() pgxmock.PgxPoolIface {
	mock, err := pgxmock.NewPool()
	if err != nil {
		log.Fatal(err)
	}

	return mock
}

func TestGetAllAnimalsRepository(t *testing.T) {
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

	tests := []struct {
		description   string
		expectedQuery string
		expectedRows  *pgxmock.Rows
		isErrExpected bool
	}{
		{
			description:   "should return no error",
			expectedQuery: `SELECT * FROM "animals" WHERE "animals"."deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
			),
			isErrExpected: false,
		},
		{
			description:   "should return error",
			expectedQuery: `SELECT * FROM "animals" WHERE"animals""deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
			),
			isErrExpected: true,
		},
		{
			description:   "should return error on scan rows",
			expectedQuery: `SELECT * FROM "animals" WHERE "animals"."deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
				"Mocked",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
				[]byte{},
			),
			isErrExpected: true,
		},
		{
			description:   "should return error on rows",
			expectedQuery: `SELECT * FROM "animals" WHERE "animals"."deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
			).RowError(1, assert.AnError),
			isErrExpected: true,
		},
	}

	for _, tc := range tests { // testCase
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			mock := mockDatabase()

			query := regexp.QuoteMeta(tc.expectedQuery)
			mock.ExpectQuery(query).WillReturnRows(tc.expectedRows)

			// Act
			animalRepository := repository.NewAnimalRepository(mock)
			animals, err := animalRepository.GetAllAnimals()
			// Assert
			if !tc.isErrExpected {
				assert.NoError(t, err)
				assert.NotNil(t, animals)
				return
			}

			assert.Error(t, err)
		})
	}
}

func TestGetAnimalByIdRepository(t *testing.T) {
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

	tests := []struct {
		description   string
		expectedQuery string
		expectedRows  *pgxmock.Rows
		isErrExpected bool
	}{
		{
			description:   "should return no error",
			expectedQuery: `SELECT * FROM "animals" WHERE "animals"."id" = $1 AND "animals"."deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
			),
			isErrExpected: false,
		},
		{
			description:   "should return error",
			expectedQuery: `SELECT * FROM "animals" WHERE"animals""deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
			),
			isErrExpected: true,
		},
		{
			description:   "should return error on scan rows",
			expectedQuery: `SELECT * FROM "animals" WHERE "animals"."id" = $1 AND "animals"."deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
				"Mocked",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
				[]byte{},
			),
			isErrExpected: true,
		},
		{
			description:   "should return error on rows",
			expectedQuery: `SELECT * FROM "animals" WHERE "animals"."id" = $1 AND "animals"."deleted_at" IS NULL`,
			expectedRows: pgxmock.NewRows([]string{
				"ID",
				"Name",
				"Breed",
				"Gender",
				"Age",
				"Size",
				"City",
				"State",
				"Dewormed",
				"Castrated",
				"Vaccinated",
				"SpecialCare",
				"Picture",
				"CreatedAt",
				"UpdatedAt",
				"DeletedAt",
			}).AddRow(
				mockAnimal.ID,
				mockAnimal.Name,
				mockAnimal.Breed,
				mockAnimal.Gender,
				mockAnimal.Age,
				mockAnimal.Size,
				mockAnimal.City,
				mockAnimal.State,
				mockAnimal.Dewormed,
				mockAnimal.Castrated,
				mockAnimal.Vaccinated,
				mockAnimal.SpecialCare,
				mockAnimal.Picture,
				mockAnimal.CreatedAt,
				mockAnimal.UpdatedAt,
				mockAnimal.DeletedAt,
			).RowError(1, assert.AnError),
			isErrExpected: true,
		},
	}

	for _, tc := range tests { // testCase
		t.Run(tc.description, func(t *testing.T) {
			// Arrange
			mock := mockDatabase()

			query := regexp.QuoteMeta(tc.expectedQuery)
			mock.ExpectQuery(query).WithArgs(mockAnimal.ID).WillReturnRows(tc.expectedRows)

			// Act
			animalRepository := repository.NewAnimalRepository(mock)
			animal, err := animalRepository.GetAnimalById(mockAnimal.ID)
			// Assert
			if !tc.isErrExpected {
				assert.NoError(t, err)
				assert.NotNil(t, animal)
				return
			}

			assert.Error(t, err)
		})
	}
}
