package handlers

import (
	"EM/internal/domain"
	"EM/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type GETPeopleHandler struct {
	useCase *usecase.ReadPersonUseCase
}

type GETPeopleResponse struct { //добавить сериалайзер
	People []domain.Person
}

func NewGETPeopleHandler(useCase *usecase.ReadPersonUseCase) *GETPeopleHandler {
	return &GETPeopleHandler{
		useCase: useCase,
	}
}

func (handler *GETPeopleHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	nationality := request.URL.Query().Get("nationality")

	offset, err := strconv.Atoi(request.URL.Query().Get("offset"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	readCommand := &usecase.ReadPersonCommand{
		Name:        name,
		Nationality: nationality,
		Offset:      offset,
		Limit:       limit,
	}

	peopleList, err := handler.useCase.ReadPersonHandler(request.Context(), readCommand)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	response := &GETPeopleResponse{
		People: peopleList,
	}
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
