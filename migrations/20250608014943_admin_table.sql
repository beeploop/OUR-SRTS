-- +goose Up
CREATE TABLE IF NOT EXISTS admin (
    id varchar(255) not null,
    fullname varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null,
    role enum('super_admin', 'staff') default 'staff',
    enabled bool default true,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY(id),
    INDEX (id, username)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS admin;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
