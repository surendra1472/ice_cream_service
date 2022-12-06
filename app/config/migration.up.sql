CREATE TABLE IF NOT EXISTS "icecream_store" (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255),
        image_closed VARCHAR(255),
        image_open VARCHAR(255),
        description VARCHAR(255),
        story VARCHAR(255),
        allergy_info VARCHAR(255),
        sourcing_values VARCHAR(255),
        ingredients VARCHAR(255),
        dietary_certifications VARCHAR(255),
        product_id VARCHAR(255));