CREATE TABLE IF NOT EXISTS terminals (
    terminal_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    station_id UUID REFERENCES stations (station_id)
)