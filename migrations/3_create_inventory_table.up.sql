CREATE TABLE inventories (
        id SERIAL PRIMARY KEY,
        organization_id INT,
        warehouse_id INT,
        user_id INT,
        document_url VARCHAR(255) DEFAULT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_organization FOREIGN KEY (organization_id) REFERENCES organizations (organization_id),
        CONSTRAINT fk_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses (warehouse_id),
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE inventory_products (
        id SERIAL PRIMARY KEY,
        inventory_id INT,
        product_id INT,
        excepted_quantity INT NOT NULL,
        factual_quantity INT NOT NULL,
        difference_quantity INT NOT NULL,
        price NUMERIC(10, 2),
        CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (product_id),
        CONSTRAINT fk_inventory FOREIGN KEY (inventory_id) REFERENCES inventories (id)
);
