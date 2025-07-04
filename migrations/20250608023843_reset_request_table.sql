-- +goose Up
CREATE TABLE IF NOT EXISTS reset_request (
    id varchar(255) not null,
    admin_id varchar(255) not null,
    expires_at timestamp not null,
    status enum('pending', 'fulfilled', 'rejected') default 'pending',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY(id),
    FOREIGN KEY(admin_id) REFERENCES admin(id),
    INDEX (id, admin_id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS reset_request;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
