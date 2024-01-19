package domain

import (
	"EM/internal/usecase/enrich_data"
	"context"
	"github.com/gofrs/uuid"
)

type Person struct {
	id          uuid.UUID
	name        string
	surname     string
	patronymic  string
	age         int
	gender      string
	nationality string
}

func (p *Person) ID() uuid.UUID       { return p.id }
func (p *Person) Name() string        { return p.name }
func (p *Person) Surname() string     { return p.surname }
func (p *Person) Patronymic() string  { return p.patronymic }
func (p *Person) Age() int            { return p.age }
func (p *Person) Gender() string      { return p.gender }
func (p *Person) Nationality() string { return p.nationality }

func NewPerson(name string, surname string, patronymic string) (*Person, error) {
	id := uuid.Must(uuid.NewV7())
	age, err := enrich_data.EnrichDataWithAge(name)
	if err != nil {
		return nil, err
	}

	gender, err := enrich_data.EnrichDataWithGender(name)
	if err != nil {
		return nil, err
	}

	nationality, err := enrich_data.EnrichDataWithNationality(name)
	if err != nil {
		return nil, err
	}

	return &Person{
		id:          id,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
		nationality: nationality,
	}, nil
}

type PersonRepository interface {
	Save(context.Context, Person) error
	Delete(context.Context, uuid.UUID) error
}
