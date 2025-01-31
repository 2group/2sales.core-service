CREATE TABLE cart (
        id SERIAL PRIMARY KEY,
        organization_id INT,
        product_id INT,
        quantity INT
);
