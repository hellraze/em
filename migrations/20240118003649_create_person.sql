-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA em;
CREATE TABLE em.person(
    person_id uuid PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    patronymic VARCHAR,
    age INT NOT NULL,
    gender VARCHAR NOT NULL,
    nationality VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA em;
DROP TABLE em.person;
-- +goose StatementEnd
