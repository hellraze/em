package domain

import (
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

func (p *Person) SetID(id uuid.UUID)                { p.id = id }
func (p *Person) SetName(name string)               { p.name = name }
func (p *Person) SetSurname(surname string)         { p.surname = surname }
func (p *Person) SetPatronymic(patronymic string)   { p.patronymic = patronymic }
func (p *Person) SetAge(age int)                    { p.age = age }
func (p *Person) SetGender(gender string)           { p.gender = gender }
func (p *Person) SetNationality(nationality string) { p.nationality = nationality }

func NewPerson(id uuid.UUID, name string, surname string, patronymic string, age int, gender string, nationality string) (*Person, error) {
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
	FindByID(context.Context, uuid.UUID) (*Person, error)
}
