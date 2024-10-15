-- +goose Up
-- +goose StatementBegin
CREATE TABLE currency (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    rate DECIMAL(20, 10) NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
