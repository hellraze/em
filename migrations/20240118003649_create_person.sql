-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA em;
CREATE TABLE em.person(
    person_id uuid PRIMARY KEY,
    name VARCHAR,
    surname VARCHAR,
    patronymic VARCHAR,
    age INT,
    gender VARCHAR,
    nationality VARCHAR

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA em;
DROP TABLE em.person;
-- +goose StatementEnd
