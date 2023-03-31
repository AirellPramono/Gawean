-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE address (
    ID          VARCHAR(100),
	Province    VARCHAR(100),
	City        VARCHAR(100),
	Descr 		TEXT

)

-- +migrate StatementEnd