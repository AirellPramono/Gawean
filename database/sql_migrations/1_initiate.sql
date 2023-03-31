-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE client (
    ID              VARCHAR(100),
	FullName        VARCHAR(256),
	Descr           TEXT,
	CreateDate      VARCHAR(100),
	UpdateDate      VARCHAR(100),
	City          VARCHAR(100),
	Province      VARCHAR(100),
	AddressID       VARCHAR(100)

)

-- +migrate StatementEnd