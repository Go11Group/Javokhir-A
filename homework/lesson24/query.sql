CREATE TABLE Cars(
    id uuid primary key not null default gen_random_uuid(),
    brand varchar not null,
    model varchar not null,
    year int
);

CREATE TABLE Users(
    id uuid primary key not null default gen_random_uuid(),
    name varchar not null,
    car_id uuid not null references Cars(id)
);

INSERT INTO Cars(id, brand, model, year) VALUES
('1e3b96df-26b8-43f9-b3b6-f1cdd9e7a8d5', 'Ford', 'Mustang', 2023),
('2a0e9e1c-3a9b-47c5-8224-7b3e9d9a72b1', 'Toyota', 'Corolla', 2022),
('2a7d1c1f-50f5-4a1b-b9a0-6c86dce7faba', 'Honda', 'Civic', 2024),
('4b9e9a2f-591e-40d4-8d92-f2a3d0a5a2f7', 'Nissan', 'Altima', 2021),
('5e6d1a7e-6f0e-4885-b8b7-1a0b1e0d8c6f', 'BMW', 'X5', 2023),
('6a0d1b9f-7a8e-41d5-b7b7-2a0b2e0c9f6a', 'Mercedes-Benz', 'C-Class', 2022),
('7a1e2b0c-8b9f-42d6-b8c7-3a1c3f0d7b7b', 'Audi', 'A4', 2023),
('8a2e3b1d-9b0f-43d7-b9d7-4a2d4f1e8c8c', 'Volkswagen', 'Golf', 2021),
('9a3e4b2e-0c1f-44d8-b0e7-5a3e5f2f9d9d', 'Tesla', 'Model 3', 2024),
('0a4e5b3f-1d2f-45d9-b1f7-6a4f6f3f0a0a', 'Hyundai', 'Elantra', 2022);

INSERT INTO Users(id, name, car_id) values
('359e2b6c-a23c-43f3-a136-ed5319513ce5', 'Javokhir Abdusamatov', '1e3b96df-26b8-43f9-b3b6-f1cdd9e7a8d5'),
('6e75148a-20c9-4efb-bc4e-03016b37c76b', 'Muhammadjon Kupalov', '7a1e2b0c-8b9f-42d6-b8c7-3a1c3f0d7b7b'),
('8e32dfae-f245-490a-973e-2c0dadab1b44', 'Saidakbar Pardaboyev', '9a3e4b2e-0c1f-44d8-b0e7-5a3e5f2f9d9d');

SELECT 
u.name, c.brand, c.model, c.year
FROM Users u
JOIN Cars c ON u.car_id = c.id;