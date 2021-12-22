CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    logo VARCHAR(155),
    banner VARCHAR(155),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE outlets (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255),
                         picture VARCHAR(155),
                         address VARCHAR(155),
                         longitude VARCHAR(155),
                         latitude VARCHAR(155),
                         brand_id INT,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                         updated_at TIMESTAMPTZ,
                         CONSTRAINT fk_brand
                             FOREIGN KEY (brand_id)
                                 REFERENCES brands(id)
);

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