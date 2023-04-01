-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    un    VARCHAR(100),
	pw	VARCHAR(100)
)

-- +migrate StatementEnd