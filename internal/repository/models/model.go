package models

import "github.com/gofrs/uuid"

type Model struct {
	PersonID    uuid.UUID
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

func NewModel() *Model {
	return &Model{}
}
