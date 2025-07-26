-- Create users table
CREATE TABLE users (
                       id BIGINT PRIMARY KEY,
                       balance BIGINT NOT NULL DEFAULT 0 CHECK ( balance >= 0 ),
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
                              amount BIGINT NOT NULL,
                              created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                              FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Seed the users table NB: This is just for simplicity, as ideally, I would not seed the users table like this.
-- CASH VALUES HAVE BEEN REPRESENTED AS SMALLEST CURRENCY UNITS TO AVOID ALL FLOATING-POINT ERRORS ie. 1025 SEEDED here represents 10.25
INSERT INTO users (id, balance, created_at, updated_at) VALUES
                                                            (1, 0, NOW(), NOW()),
                                                            (2, 1025, NOW(), NOW()),
                                                            (3, 10, NOW(), NOW());
