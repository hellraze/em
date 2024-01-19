package usecase

import (
	"EM/internal/domain"
	"context"
	"github.com/gofrs/uuid"
)

type DeletePersonUseCase struct {
	personRepository domain.PersonRepository
}

func NewDeletePersonUseCase(personRepository domain.PersonRepository) *DeletePersonUseCase {
	return &DeletePersonUseCase{
		personRepository: personRepository,
	}
}

type DeletePersonCommand struct {
	ID uuid.UUID
}

func (useCase *DeletePersonUseCase) DeleteUserHandler(ctx context.Context, command *DeletePersonCommand) error {
	err := useCase.personRepository.Delete(ctx, command.ID)
	return err
}
