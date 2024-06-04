CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    username VARCHAR NOT NULL ,
    email VARCHAR NOT NULL ,
    password VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    phone_number VARCHAR,
    address TEXT,
    profile_picture_url TEXT,
    roles VARCHAR DEFAULT 'user',
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN DEFAULT TRUE,
    CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);
CREATE INDEX user_id_idx ON users (id, username, first_name, last_name);
DROP INDEX user_id_idx;

CREATE INDEX user_id_idx_hash ON USING HASH(id)

TRUNCATE users;
DROP TABLE users;
EXPLAIN(ANALYZE)
SELECT *FROM users WHERE id = '877f37dd-d22b-4111-8cb1-ce68abf77683';

SELECT *from users OFFSET 500000 LIMIT 100;