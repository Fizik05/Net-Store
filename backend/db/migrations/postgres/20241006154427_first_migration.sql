-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Products (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR,
    price INT NOT NULL,
    image_url VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Users (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd
