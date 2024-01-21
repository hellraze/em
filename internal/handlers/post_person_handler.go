package handlers

import (
	"EM/internal/pkg/logging"
	"EM/internal/usecase"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type POSTPersonHandler struct {
	useCase *usecase.CreatePersonUseCase
}

type POSTPersonRequest struct {
	Name       string
	Surname    string
	Patronymic string
}

type POSTPersonResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewPOSTPersonHandler(useCase *usecase.CreatePersonUseCase) *POSTPersonHandler {
	return &POSTPersonHandler{
		useCase: useCase,
	}
}

func (handler *POSTPersonHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log := logging.NewLog()
	log.Init()

	var body POSTPersonRequest
	ctx := request.Context()
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	command := &usecase.CreatePersonCommand{
		Name:       body.Name,
		Surname:    body.Surname,
		Patronymic: body.Patronymic,
	}

	person, err := handler.useCase.CreateUserHandler(ctx, command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Log.WithFields(logrus.Fields{
		"id":   person.ID(),
		"name": person.Name(),
	}).Info("Получен запрос на добавление пользователя")

	response := &POSTPersonResponse{
		ID: person.ID(),
	}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
