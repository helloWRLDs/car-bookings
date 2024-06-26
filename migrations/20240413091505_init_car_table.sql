-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    vendor VARCHAR(30) NOT NULL,
    model VARCHAR(60) NOT NULL,
    year SMALLINT NOT NULL,
    is_available BOOLEAN,
    body VARCHAR(30) NOT NULL,
    engine_capacity FLOAT NOT NULL,
    mileage INT NOT NULL,
    color VARCHAR(30) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
-- +goose StatementEnd
