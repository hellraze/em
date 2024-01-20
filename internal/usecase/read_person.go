package usecase

import (
	"EM/internal/domain"
	"context"
	"github.com/gofrs/uuid"
)

type ReadPersonUseCase struct {
	personRepository domain.PersonRepository
}

func NewReadPersonUseCase(personRepository domain.PersonRepository) *ReadPersonUseCase {
	return &ReadPersonUseCase{
		personRepository: personRepository,
	}
}

type ReadPersonCommand struct {
	ID          uuid.UUID
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

func (useCase *ReadPersonUseCase) ReadUserHandler(ctx context.Context, command *ReadPersonCommand) (*domain.Person, error) {
	person, err := domain.NewPerson(command.ID, command.Name, command.Surname, command.Patronymic, command.Age, command.Gender, command.Nationality)
	err = useCase.personRepository.Update(ctx, command.ID, command.Name, command.Surname, command.Patronymic, command.Age, command.Gender, command.Nationality)
	return person, err
}
