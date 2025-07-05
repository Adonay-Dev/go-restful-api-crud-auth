-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_tbl (
    id INT NOT NULL PRIMARY KEY,
    email VARCHAR NOT NULL,
    fullName VARCHAR NOT NULL,
    PASSWORD VARCHAR NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd