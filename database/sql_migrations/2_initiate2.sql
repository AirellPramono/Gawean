-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE gawean (
    ID          INT,
	ClientID    VARCHAR(100),
	ClientName  VARCHAR(256),
	Descr       TEXT,
	CreateDate  VARCHAR(100),
	UpdateDate  VARCHAR(100),
	CategoryID  INT

)

-- +migrate StatementEnd