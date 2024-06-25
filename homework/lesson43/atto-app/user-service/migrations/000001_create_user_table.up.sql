CREATE TABLE IF NOT EXISTS users(
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR,
    last_name VARCHAR,
    email VARCHAR,
    age INT CHECK(age > 18),
    phone VARCHAR,
    role VARCHAR DEFAULT 'user',
    password_hash VARCHAR,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    delete_at TIMESTAMP DEFAULT NULL
);