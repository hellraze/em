package usecase

import (
	"EM/internal/domain"
	"context"
)

type CreatePersonUseCase struct {
	personRepository domain.PersonRepository
}

func NewCreatePersonUseCase(personRepository domain.PersonRepository) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		personRepository: personRepository,
	}
}

type CreatePersonCommand struct {
	Name       string
	Surname    string
	Patronymic string
}

func (useCase *CreatePersonUseCase) CreateUserHandler(ctx context.Context, command *CreatePersonCommand) (*domain.Person, error) {
	person, err := domain.NewPerson(command.Name, command.Surname, command.Patronymic)
	err = useCase.personRepository.Save(ctx, *person)
	return person, err
}
