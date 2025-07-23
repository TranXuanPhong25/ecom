CREATE TABLE IF NOT EXISTS brands
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255)                NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS products
(
    id          BIGSERIAL PRIMARY KEY,
    shop_id  uuid NOT NULL,
    name        VARCHAR(255)                NOT NULL UNIQUE,
    description TEXT,
    category_id BIGINT,
    brand_id BIGINT REFERENCES brands (id),
    is_active   BOOLEAN                              DEFAULT TRUE,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS product_variants
(
    id             BIGSERIAL PRIMARY KEY,
    product_id     BIGSERIAL                   NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    price          DECIMAL(10, 2)              NOT NULL,
    attributes     JSONB,
    is_active      BOOLEAN                              DEFAULT TRUE,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    stock_quantity INT                         NOT NULL

);
CREATE TABLE IF NOT EXISTS product_variant_skus
(
    id         BIGSERIAL PRIMARY KEY,
    variant_id BIGSERIAL                   NOT NULL REFERENCES product_variants (id) ON DELETE CASCADE,
    sku        VARCHAR(100)                NOT NULL UNIQUE,

    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS variant_images
(
    id         BIGSERIAL PRIMARY KEY,
    variant_id BIGSERIAL                   NOT NULL REFERENCES product_variants (id) ON DELETE CASCADE,
    image_url  TEXT                        NOT NULL,
    is_primary BOOLEAN                              DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
