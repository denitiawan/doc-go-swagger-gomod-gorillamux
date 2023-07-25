-- +migrate Up
-- +migrate StatementBegin

insert into `users` (`name`,`username`,`password`) values ('demo','demo','demo');

-- +migrate StatementEnd