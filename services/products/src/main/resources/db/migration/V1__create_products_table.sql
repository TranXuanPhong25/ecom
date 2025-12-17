CREATE TYPE product_status AS ENUM ('DRAFT', 'ACTIVE', 'ARCHIVED');
CREATE TABLE IF NOT EXISTS brands
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS products
(
    id              BIGSERIAL PRIMARY KEY,
    shop_id         uuid           NOT NULL,
    name            VARCHAR(255)   NOT NULL UNIQUE,
    description     TEXT,
    cover_image TEXT           NOT NULL,
    specs           JSONB,
    images          TEXT ARRAY,
    category_id     BIGINT,
    brand_id        BIGINT REFERENCES brands (id),
    status          product_status NOT NULL DEFAULT 'DRAFT',
    created_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS product_variants
(
    id             BIGSERIAL PRIMARY KEY,
    product_id     BIGSERIAL      NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    price          DECIMAL(10, 2) NOT NULL,
    attributes     JSONB,
    images         TEXT ARRAY,
    sku            VARCHAR(255) UNIQUE,
    is_active      BOOLEAN                 DEFAULT TRUE,
    created_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    stock_quantity INT            NOT NULL
);
