package repository

import (
	"EM/internal/domain"
	"EM/internal/pkg/logging"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type PersonRepository struct {
	pool *pgxpool.Pool
}

func NewPersonRepository(pool *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{
		pool: pool,
	}
}

type Model struct {
	PersonID    uuid.UUID
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

func NewModel() *Model {
	return &Model{}
}

func (personRepository *PersonRepository) Save(ctx context.Context, person domain.Person) error {
	log := logging.NewLog()
	log.Init()

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
	INSERT INTO EM.person(person_id, name, surname, patronymic, age, gender, nationality) 
	VALUES(@id, @name, @surname, @patronymic, @age, @gender, @nationality)`, args)

	log.Log.WithFields(logrus.Fields{
		"id":   person.ID(),
		"name": person.Name(),
	}).Info("Новый пользователь добавлен")

	return err
}

func (personRepository *PersonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	log := logging.NewLog()
	log.Init()

	_, err := personRepository.pool.Exec(ctx, "DELETE FROM EM.person WHERE person_id=$1", id) //логи
	if err != nil {
		return err
	}

	log.Log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Пользователь с данным id удален")

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

func (personRepository *PersonRepository) Update(ctx context.Context, id uuid.UUID, name string, surname string, patronymic string, age int, gender string, nationality string) error {
	log := logging.NewLog()
	log.Init()

	person, err := personRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if name != "" {
		person.SetName(name)
	}
	if surname != "" {
		person.SetSurname(surname)
	}
	if patronymic != "" {
		person.SetPatronymic(patronymic)
	}
	if age != 0 {
		person.SetAge(age)
	}
	if gender != "" {
		person.SetGender(gender)
	}
	if nationality != "" {
		person.SetNationality(nationality)
	}
	args := pgx.NamedArgs{
		"id":          person.ID(),
		"name":        person.Name(),
		"surname":     person.Surname(),
		"patronymic":  person.Patronymic(),
		"age":         person.Age(),
		"gender":      person.Gender(),
		"nationality": person.Nationality(),
	}

	_, err = personRepository.pool.Exec(ctx, `
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

	log.Log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Пользователь с данным id изменен")

	if err != nil {
		return err
	}
	return nil
}

func (personRepository *PersonRepository) Read(ctx context.Context, nameFilter string, nationalityFilter string, offset int, limit int) ([]domain.Person, error) {
	log := logging.NewLog()
	log.Init()

	limitUint64 := uint64(limit)
	offsetUint64 := uint64(offset)

	var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query := psql.Select("*").From("EM.person")
	if nameFilter != "" {
		query = query.Where(squirrel.Eq{"name": nameFilter})
	}
	if nationalityFilter != "" {
		query = query.Where(squirrel.Eq{"nationality": nationalityFilter})
	}
	query = query.Offset(offsetUint64)
	query = query.Limit(limitUint64)
	sql, _, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := personRepository.pool.Query(ctx, sql, nameFilter, nationalityFilter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []domain.Person

	for rows.Next() {
		model := NewModel()

		err = rows.Scan(
			&model.PersonID,
			&model.Name,
			&model.Surname,
			&model.Patronymic,
			&model.Age,
			&model.Gender,
			&model.Nationality,
		)
		person, err := domain.NewPerson(model.PersonID, model.Name, model.Surname, model.Patronymic, model.Age, model.Gender, model.Nationality)
		if err != nil {
			return nil, err
		}
		people = append(people, *person)

		log.Log.WithFields(logrus.Fields{
			"id": person.ID(),
		}).Info("Пользователь с данным id найден")
	}

	return people, nil
}
