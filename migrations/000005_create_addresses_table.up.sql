-- 005_create_addresses_table.sql
-- Delivery addresses for users

CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipient_name VARCHAR(191) NOT NULL,
    phone_number VARCHAR(32) NOT NULL,
    street_address TEXT NOT NULL,
    city VARCHAR(191),
    province VARCHAR(191),
    postal_code VARCHAR(32),
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
