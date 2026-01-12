-- +goose Up
CREATE TABLE crypto(
        id          UUID PRIMARY KEY,
	created_at  TIMESTAMP NOT NULL,
	updated_at  TIMESTAMP NOT NULL,
	crypto_key  TEXT NOT NULL,
	crypto_list JSONB NOT NULL
);

-- +goose Down
DROP TABLE crypto;
