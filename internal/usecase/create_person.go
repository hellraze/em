package usecase

import (
	"EM/internal/domain"
	enrich_data2 "EM/internal/services/enrich_data"
	"context"
	"github.com/gofrs/uuid"
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
	id := uuid.Must(uuid.NewV7())
	age, err := enrich_data2.EnrichDataWithAge(command.Name) //перенести в юзкейс
	if err != nil {
		return nil, err
	}

	gender, err := enrich_data2.EnrichDataWithGender(command.Name)
	if err != nil {
		return nil, err
	}

	nationality, err := enrich_data2.EnrichDataWithNationality(command.Name)
	if err != nil {
		return nil, err
	}

	person, err := domain.NewPerson(id, command.Name, command.Surname, command.Patronymic, age, gender, nationality)
	err = useCase.personRepository.Save(ctx, *person)
	return person, err
}
