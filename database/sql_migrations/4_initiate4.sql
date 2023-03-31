-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    ID          VARCHAR(100),
	Fullname	VARCHAR(256)
)

-- +migrate StatementEnd