DROP TABLE cart;

CREATE TABLE carts (
        id SERIAL PRIMARY KEY,
        organization_id INT
);

CREATE TABLE cart_products (
        id SERIAL PRIMARY KEY,
        cart_id INT,
        product_id INT,
        quantity INT
);

