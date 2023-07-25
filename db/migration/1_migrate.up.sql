-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users
(
    id       BIGINT NOT NULL AUTO_INCREMENT,
    name     VARCHAR(256),
    username VARCHAR(256),
    password VARCHAR(256),
    primary key (id)
);

-- +migrate StatementEnd