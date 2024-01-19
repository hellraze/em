package handlers

import (
	"EM/internal/usecase"
	"encoding/json"
	"github.com/gofrs/uuid"
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

	id, err := uuid.FromString(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "ID parameter is required", http.StatusBadRequest)
		return
	}
	if id == uuid.Nil {
		http.Error(writer, "ID parameter is required", http.StatusBadRequest)
		return
	}

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
