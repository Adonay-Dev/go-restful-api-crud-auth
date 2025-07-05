-- +goose Up
-- +goose StatementBegin
CREATE TABLE event_tbl (
    id INT NOT NULL PRIMARY KEY,
    ownerId INT NOT NULL,
    eventName VARCHAR NOT NULL,
    eventDescription VARCHAR NOT NULL,
    date DATE NOT NULL,
    location VARCHAR NOT NULL,
    FOREIGN KEY (ownerId) REFERENCES user_tbl (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd