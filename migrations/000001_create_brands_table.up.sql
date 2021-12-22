CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    logo VARCHAR(155),
    banner VARCHAR(155),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);