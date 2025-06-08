-- +goose Up
CREATE TABLE IF NOT EXISTS program (
    id varchar(255) not null,
    title varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY(id),
    INDEX (id, title)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS program;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
