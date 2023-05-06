-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stores(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    availability BOOLEAN NOT NULL DEFAULT TRUE
);
CREATE TABLE IF NOT EXISTS items(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    size INTEGER
);
CREATE TABLE IF NOT EXISTS store_availability(
    store_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    item_quantity INTEGER NOT NULL DEFAULT 0,
    reserved_item INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE store_availability;
DROP TABLE stores;
DROP TABLE items;
-- +goose StatementEnd
