-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          serial NOT NULL PRIMARY KEY,
	name        varchar(255) NOT NULL,
	hobby        varchar(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
