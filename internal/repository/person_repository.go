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

func (personRepository *PersonRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Person, error) {
	var (
		name        string
		surname     string
		patronymic  string
		age         int
		gender      string
		nationality string
	)
	err := personRepository.pool.QueryRow(ctx, "SELECT * FROM EM.person WHERE person_id=$1", id).Scan(&id, &name, &surname, &patronymic, &age, &gender, &nationality)
	if err != nil {
		return nil, err
	}
	person, err := domain.NewPerson(id, name, surname, patronymic, age, gender, nationality)
	return person, nil
}

func (personRepository *PersonRepository) UpdatePerson(ctx context.Context, person domain.Person) error {
	args := pgx.NamedArgs{
		"id":          person.ID(),
		"name":        person.Name(),
		"surname":     person.Surname(),
		"patronymic":  person.Patronymic(),
		"age":         person.Age(),
		"gender":      person.Gender(),
		"nationality": person.Nationality(),
	}

	_, err := personRepository.pool.Exec(ctx, `
		UPDATE EM.person 
		SET 
			name = @name,
			surname = @surname,
			patronymic = @patronymic,
			age = @age,
			gender = @gender,
			nationality = @nationality
		WHERE 
			person_id = @id
	`, args)

	if err != nil {
		return err
	}

	return nil
}

func (personRepository *PersonRepository) Update(ctx context.Context, id uuid.UUID, name string, surname string, patronymic string, age int, gender string, nationality string) error {
	person, err := personRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	switch {
	case name != "":
		person.SetName(name)
	case surname != "":
		person.SetSurname(surname)
	case patronymic != "":
		person.SetPatronymic(patronymic)
	case age != 0:
		person.SetAge(age)
	case gender != "":
		person.SetGender(gender)
	case nationality != "":
		person.SetNationality(nationality)
	}
	err = personRepository.UpdatePerson(ctx, *person)
	if err != nil {
		return err
	}
	return nil
}
