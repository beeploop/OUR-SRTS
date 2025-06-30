-- +goose Up
CREATE TABLE IF NOT EXISTS envelope (
    id varchar(255) not null,
    owner varchar(255) not null unique,
    location varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY(id),
    INDEX(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS envelope;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
