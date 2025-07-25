-- Create users table
CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00 CHECK ( balance >= 0.00 ),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
); 

-- Create transactions table
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    transaction_id VARCHAR(255) UNIQUE NOT NULL,
    source_type VARCHAR(50) NOT NULL,
    state VARCHAR(10) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);