CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    points INT DEFAULT 0 NOT NULL,
    balance INT DEFAULT 0 NOT NULL
);

CREATE TABLE brands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    price INT DEFAULT 0 NOT NULL,
    brand_id UUID REFERENCES brands(id) ON DELETE CASCADE
);

CREATE TABLE vouchers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(255) UNIQUE NOT NULL,
    cost_in_point INT DEFAULT 0 NOT NULL,
    expiration TIMESTAMP NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('discount', 'cashback')), 
    value INT DEFAULT 0 NOT NULL,
    brand_id UUID REFERENCES brands(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    voucher_id UUID REFERENCES vouchers(id) ON DELETE SET NULL,
    total INT DEFAULT 0 NOT NULL
);
