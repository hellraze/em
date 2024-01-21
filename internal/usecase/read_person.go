package usecase

import (
	"EM/internal/domain"
	"context"
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
	Name        string
	Nationality string
	Offset      int
	Limit       int
}

func (useCase *ReadPersonUseCase) ReadPersonHandler(ctx context.Context, command *ReadPersonCommand) ([]domain.Person, error) {
	people, err := useCase.personRepository.Read(ctx, command.Name, command.Nationality, command.Offset, command.Limit)
	return people, err
}
