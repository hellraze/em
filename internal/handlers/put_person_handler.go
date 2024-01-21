package handlers

import (
	"EM/internal/pkg/logging"
	"EM/internal/usecase"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PutPersonHandler struct {
	useCase *usecase.PutPersonUseCase
}

type PutPersonRequest struct {
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

type PutPersonResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewPutPersonHandler(useCase *usecase.PutPersonUseCase) *PutPersonHandler {
	return &PutPersonHandler{
		useCase: useCase,
	}
}

func (handler *PutPersonHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log := logging.NewLog()
	log.Init()

	id, err := uuid.FromString(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "ID parameter is required", http.StatusBadRequest)
		return
	}
	if id == uuid.Nil {
		http.Error(writer, "ID parameter is required", http.StatusBadRequest)
		return
	}
	var body PutPersonRequest
	ctx := request.Context()

	err = json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	command := &usecase.PutPersonCommand{
		ID:          id,
		Name:        body.Name,
		Surname:     body.Surname,
		Patronymic:  body.Patronymic,
		Age:         body.Age,
		Gender:      body.Gender,
		Nationality: body.Nationality,
	}

	log.Log.WithFields(logrus.Fields{
		"id": command.ID,
	}).Info("Получен запрос на изменение пользователя с данным id")

	person, err := handler.useCase.PutUserHandler(ctx, command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &POSTPersonResponse{
		ID: person.ID(),
	}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
