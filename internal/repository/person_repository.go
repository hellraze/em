package repository

import (
	"EM/internal/domain"
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonRepository struct {
	pool *pgxpool.Pool
}

func NewPersonRepository(pool *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{
		pool: pool,
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
	_, err := personRepository.pool.Exec(ctx, "INSERT INTO EM.person(person_id, name, surname, patronymic, age, gender, nationality) VALUES(@id, @name, @surname, @patronymic, @age, @gender, @nationality)", args)
	return err
}

func (personRepository *PersonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := personRepository.pool.Exec(ctx, "DELETE FROM EM.person WHERE person_id=$1", id) //логи
	if err != nil {
		return err
	}
	return nil
}
