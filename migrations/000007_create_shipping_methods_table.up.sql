-- 007_create_shipping_methods_table.sql
-- Shipping method definitions with cost and estimated delivery days

CREATE TABLE shipping_methods (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(191) NOT NULL,
    description TEXT,
    cost DECIMAL(12,2) NOT NULL DEFAULT 0.00,
    estimated_days INTEGER,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
