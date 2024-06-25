CREATE TABLE IF NOT EXISTS stations(
    station_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    address VARCHAR DEFAULT 'Tashkent'
)