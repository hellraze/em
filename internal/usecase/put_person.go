package usecase

import (
	"EM/internal/domain"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
)

type PutPersonUseCase struct {
	personRepository domain.PersonRepository
}

func NewPutPersonUseCase(personRepository domain.PersonRepository) *PutPersonUseCase {
	return &PutPersonUseCase{
		personRepository: personRepository,
	}
}

type PutPersonCommand struct {
	ID          uuid.UUID
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

func (useCase *PutPersonUseCase) PutUserHandler(ctx context.Context, command *PutPersonCommand) (*domain.Person, error) {
	person, err := domain.NewPerson(command.ID, command.Name, command.Surname, command.Patronymic, command.Age, command.Gender, command.Nationality)
	id, err := useCase.personRepository.Update(ctx, command.ID, command.Name, command.Surname, command.Patronymic, command.Age, command.Gender, command.Nationality)
	fmt.Println(id)
	return person, err
}
