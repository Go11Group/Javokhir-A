CREATE TABLE IF NOT EXISTS weather_report (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    location VARCHAR(255) NOT NULL,
    temperature FLOAT NOT NULL,
    humidity FLOAT NOT NULL,
    pressure FLOAT NOT NULL,
    wind_speed FLOAT NOT NULL, 
    reported_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
