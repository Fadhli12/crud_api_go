CREATE TABLE products
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255),
    picture    VARCHAR(155),
    price      INT,
    brand_id   INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_brand
        FOREIGN KEY (brand_id)
            REFERENCES brands (id)
);