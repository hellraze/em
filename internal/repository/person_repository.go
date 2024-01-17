package repository

import (
	"EM/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonRepository struct {
	Pool *pgxpool.Pool
}

func NewPersonRepository(pool *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{
		Pool: pool,
	}
}

func (personRepository *PersonRepository) Save(ctx context.Context, person domain.Person) error {
	args := pgx.NamedArgs{
		"id":          person.ID(),
		"name":        person.Name(),
		"surname":     person.Surname(),
		"patronymic":  person.Patronymic(),
		"age":         person.Age(),
		"gender":      person.Gender(),
		"nationality": person.Nationality(),
	}
	_, err := personRepository.Pool.Exec(ctx, "INSERT INTO EM.person(person_id, name, surname, patronmic, age, gender, nationality) VALUES(@id, @name, @surname, @patronymic, @age, @gender, @nationality)", args)
	return err
}
