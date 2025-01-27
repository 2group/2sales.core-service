CREATE TABLE movings (
    id SERIAL PRIMARY KEY,
    organization_id INT,
    from_warehouse_id INT,
    to_warehouse_id INT,
    user_id INT,
    document_url VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_organization FOREIGN KEY (organization_id) REFERENCES organizations (organization_id),
    CONSTRAINT fk_to_warehouse FOREIGN KEY (to_warehouse_id) REFERENCES warehouses (warehouse_id),
    CONSTRAINT fk_from_warehouse FOREIGN KEY (from_warehouse_id) REFERENCES warehouses (warehouse_id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE moving_products (
    id SERIAL PRIMARY KEY,
    moving_id INT,
    product_id INT,
    quantity INT NOT NULL CHECK (quantity >= 0),
    price NUMERIC(10, 2),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (product_id),
    CONSTRAINT fk_moving FOREIGN KEY (moving_id) REFERENCES movings (id)
)
