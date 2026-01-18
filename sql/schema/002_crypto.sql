-- +goose Up
CREATE INDEX crypto_search_index ON crypto (crypto_key);

-- +goose Down
DROP INDEX crypto_search_index;
