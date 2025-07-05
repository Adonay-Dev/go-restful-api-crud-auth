-- +goose Up
-- +goose StatementBegin
CREATE TABLE attendees_tbl (
    id INT NOT NULL PRIMARY KEY,
    evenId INT NOT NULL,
    userId INT NOT NULL,
    FOREIGN KEY (userId) REFERENCES user_tbl (id),
    FOREIGN KEY (evenId) REFERENCES event_tbl (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd