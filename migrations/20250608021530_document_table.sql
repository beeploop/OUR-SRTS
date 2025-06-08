-- +goose Up
CREATE TABLE IF NOT EXISTS document (
    id varchar(255) not null,
    type_id varchar(255) not null,
    filename varchar(255) not null,
    storage_path varchar(255) not null,
    uploaded_at timestamp default current_timestamp,
    PRIMARY KEY(id),
    FOREIGN KEY(type_id) REFERENCES document_type(id),
    INDEX (id, type_id, filename)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS document;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
