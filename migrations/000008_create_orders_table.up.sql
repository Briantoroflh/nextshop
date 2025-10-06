-- 008_create_orders_table.sql
-- Main orders with payment, shipping, and status

CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipped', 'delivered', 'cancelled');
CREATE TYPE payment_status AS ENUM ('pending', 'success', 'failed', 'refunded');

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    buyer_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    address_id BIGINT REFERENCES addresses(id) ON DELETE SET NULL,
    payment_method_id BIGINT REFERENCES payment_methods(id) ON DELETE SET NULL,
    shipping_method_id BIGINT REFERENCES shipping_methods(id) ON DELETE SET NULL,
    total_amount DECIMAL(12,2) NOT NULL,
    order_status order_status NOT NULL DEFAULT 'pending',
    payment_status payment_status NOT NULL DEFAULT 'pending',
    note TEXT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
