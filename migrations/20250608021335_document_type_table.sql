-- +goose Up
CREATE TABLE IF NOT EXISTS document_type (
    id varchar(255) not null,
    title varchar(255) not null unique,
    full_title varchar(255) not null,
    is_stable boolean default true,
    allow_multiple boolean default false,
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
