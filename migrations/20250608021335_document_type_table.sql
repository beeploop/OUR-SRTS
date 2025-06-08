-- +goose Up
CREATE TABLE IF NOT EXISTS document_type (
    id varchar(255) not null,
    name varchar(255) not null unique,
    PRIMARY KEY(id),
    INDEX (id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS document_type;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
