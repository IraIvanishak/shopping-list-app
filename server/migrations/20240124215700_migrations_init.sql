-- +goose Up
-- +goose StatementBegin
CREATE TABLE shopping_list (
    id SERIAL PRIMARY KEY,
	name VARCHAR (255) NOT NULL,
	amount INT NOT NULL
);

INSERT INTO shopping_list (name, amount) VALUES 
('Apples', 5),
('Oranges', 3),
('Milk', 2),
('Bread', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shopping_list;
-- +goose StatementEnd
