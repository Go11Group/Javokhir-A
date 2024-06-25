CREATE TABLE if NOT EXISTS cards (
    card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_number VARCHAR(255) NOT NULL,
    card_type VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    created_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP DEFAULT NULL
)