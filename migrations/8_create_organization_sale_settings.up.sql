CREATE TABLE organization_sale_settings (
       id SERIAL PRIMARY KEY,
       organization_id INT,
       category_id VARCHAR(255),
       profit_percent NUMERIC(10, 2)
);

CREATE TABLE sale_releated_expences (
       id SERIAL PRIMARY KEY,
       organization_sale_setting_id INT,
       name VARCHAR(255),
       value NUMERIC(10, 2),
       FOREIGN KEY (organization_sale_setting_id) REFERENCES organization_sale_settings(id)
);
