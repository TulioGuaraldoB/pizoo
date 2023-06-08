package repository

import (
	"context"
	"time"

	"github.com/TulioGuaraldoB/pizoo/internal/domain/entity"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/postgres"
)

type animalRepository struct {
	context context.Context
	db      postgres.IDbPool
}

func NewAnimalRepository(db postgres.IDbPool) *animalRepository {
	return &animalRepository{context: context.Background(), db: db}
}

func (r *animalRepository) GetAllAnimals() ([]entity.Animal, error) {
	animals := make([]entity.Animal, 0)
	rows, err := r.db.Query(r.context, `SELECT * FROM "animals" WHERE "animals"."deleted_at" IS NULL`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		animal := new(entity.Animal)
		if err := rows.Scan(
			&animal.ID,
			&animal.Name,
			&animal.Breed,
			&animal.Gender,
			&animal.Age,
			&animal.Size,
			&animal.City,
			&animal.State,
			&animal.Dewormed,
			&animal.Castrated,
			&animal.Vaccinated,
			&animal.SpecialCare,
			&animal.Picture,
			&animal.CreatedAt,
			&animal.UpdatedAt,
			&animal.DeletedAt,
		); err != nil {
			return nil, err
		}

		animals = append(animals, *animal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return animals, nil
}

func (r *animalRepository) GetAnimalById(animalId uint) (*entity.Animal, error) {
	animal := new(entity.Animal)
	rows, err := r.db.Query(r.context, `SELECT * FROM "animals" WHERE "animals"."id" = $1 AND "animals"."deleted_at" IS NULL`, animalId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&animal.ID,
			&animal.Name,
			&animal.Breed,
			&animal.Gender,
			&animal.Age,
			&animal.Size,
			&animal.City,
			&animal.State,
			&animal.Dewormed,
			&animal.Castrated,
			&animal.Vaccinated,
			&animal.SpecialCare,
			&animal.Picture,
			&animal.CreatedAt,
			&animal.UpdatedAt,
			&animal.DeletedAt,
		); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return animal, nil
}

func (r *animalRepository) CreateAnimal(animal *entity.Animal) error {
	if _, err := r.db.Exec(r.context, `INSERT INTO "animals" (animal, breed, gender, age, size, city, state, dewormed, castrated, vaccinated, special_care, picture, created_at, udpated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`); err != nil {
		return err
	}

	return nil
}

func (r *animalRepository) DeleteAnimal(animalId uint) error {
	_, err := r.db.Exec(r.context, `UPDATE "animals" SET "animals"."deleted_at" = $1 AND "animals"."id" = $2`, time.Now(), animalId)
	if err != nil {
		return err
	}

	return nil
}
