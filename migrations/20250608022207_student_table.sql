-- +goose Up
CREATE TABLE IF NOT EXISTS student (
    control_number varchar(255) not null,
    first_name varchar(255) not null,
    middle_name varchar(255),
    last_name varchar(255) not null,
    suffix varchar(255),
    student_type enum('non_transferee', 'transferee', 'graduate') default 'non_transferee',
    civil_status enum('single', 'married') default 'single',
    program_id varchar(255) not null,
    major_id varchar(255) not null,
    envelope_id varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY(control_number),
    FOREIGN KEY(envelope_id) REFERENCES envelope(id),
    INDEX (control_number, first_name, last_name)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS student;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
