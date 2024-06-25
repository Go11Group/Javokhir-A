CREATE TYPE trans_type AS ENUM ('credit', 'debit');

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_id UUID REFERENCES cards(card_id),
    amount FLOAT,
    transaction_type trans_type,
    terminal_id UUID REFERENCES terminals(terminal_id) DEFAULT NULL,
    transaction_date TIMESTAMP DEFAULT current_timestamp
);
