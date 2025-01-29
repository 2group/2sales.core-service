CREATE TABLE characteristic_values (
       id SERIAL PRIMARY KEY,
       characteristic_id INT REFERENCES product_characteristics(id),
       text_value VARCHAR(255) NULL,
       num_value NUMERIC(10, 2) NULL
);
