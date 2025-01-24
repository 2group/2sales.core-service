CREATE TABLE write_offs (
    id SERIAL PRIMARY KEY,
    organization_id INT NOT NULL,
    warehouse_id INT NOT NULL,
    document_url VARCHAR(255) DEFAULT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_organization FOREIGN KEY (organization_id) REFERENCES organizations (organization_id),
    CONSTRAINT fk_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses (warehouse_id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE write_off_products (
    id SERIAL PRIMARY KEY,
    write_off_id INT,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    price NUMERIC(10, 2),
    reason VARCHAR(255),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (product_id),
    CONSTRAINT fk_write_off FOREIGN KEY (write_off_id) REFERENCES write_offs (id)
);
