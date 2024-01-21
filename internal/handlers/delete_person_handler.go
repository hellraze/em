package handlers

import (
	"EM/internal/pkg/logging"
	"EM/internal/usecase"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type DeletePersonHandler struct {
	useCase *usecase.DeletePersonUseCase
}

type DeletePersonRequest struct {
	ID uuid.UUID
}

type DeletePersonResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewDeletePersonHandler(useCase *usecase.DeletePersonUseCase) *DeletePersonHandler {
	return &DeletePersonHandler{
		useCase: useCase,
	}
}
func (handler *DeletePersonHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

	log.Log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Получен запрос на удаление пользователя с данным id")

	ctx := request.Context()

	command := &usecase.DeletePersonCommand{
		ID: id,
	}

	err = handler.useCase.DeleteUserHandler(ctx, command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &POSTPersonResponse{
		ID: id,
	}

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
