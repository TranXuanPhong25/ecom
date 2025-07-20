CREATE TABLE products
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255)                NOT NULL UNIQUE,
    description TEXT,
    category_id BIGINT,
    brand_id    BIGINT,
    is_active   BOOLEAN                              DEFAULT TRUE,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE product_variants
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
CREATE TABLE product_variant_skus
(
    id         BIGSERIAL PRIMARY KEY,
    variant_id BIGSERIAL                   NOT NULL REFERENCES product_variants (id) ON DELETE CASCADE,
    sku        VARCHAR(100)                NOT NULL UNIQUE,

    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE variant_images
(
    id         BIGSERIAL PRIMARY KEY,
    variant_id BIGSERIAL                   NOT NULL REFERENCES product_variants (id) ON DELETE CASCADE,
    image_url  TEXT                        NOT NULL,
    is_primary BOOLEAN                              DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);