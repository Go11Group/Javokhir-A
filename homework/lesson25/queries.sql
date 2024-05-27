
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE cars (
    car_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_name VARCHAR(100) NOT NULL
);

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_name VARCHAR(100) NOT NULL
);

CREATE TABLE car_owner (
    car_id UUID,
    user_id UUID,
    PRIMARY KEY (car_id, user_id),
    FOREIGN KEY (car_id) REFERENCES cars (car_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE cars (
    car_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_name VARCHAR(100) NOT NULL
);

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_name VARCHAR(100) NOT NULL
);

CREATE TABLE car_owner (
    car_id UUID,
    user_id UUID,
    PRIMARY KEY (car_id, user_id),
    FOREIGN KEY (car_id) REFERENCES cars (car_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);


INSERT INTO car_owner (car_id, user_id) 
SELECT car_id, user_id 
FROM (SELECT car_id FROM cars ORDER BY random() LIMIT 10) AS random_cars,
     (SELECT user_id FROM users ORDER BY random() LIMIT 10) AS random_users
LIMIT 10;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE cars (
    car_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    car_name VARCHAR(100) NOT NULL
);

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_name VARCHAR(100) NOT NULL
);

CREATE TABLE car_owner (
    car_id UUID,
    user_id UUID,
    PRIMARY KEY (car_id, user_id),
    FOREIGN KEY (car_id) REFERENCES cars (car_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO cars (car_name) VALUES
('Car A'),
('Car B'),
('Car C'),
('Car D'),
('Car E'),
('Car F'),
('Car G'),
('Car H'),
('Car I'),
('Car J');


INSERT INTO users (user_name) VALUES
('User 1'),
('User 2'),
('User 3'),
('User 4'),
('User 5'),
('User 6'),
('User 7'),
('User 8'),
('User 9'),
('User 10');


INSERT INTO car_owner (car_id, user_id) 
SELECT car_id, user_id 
FROM (SELECT car_id FROM cars ORDER BY random() LIMIT 10) AS random_cars,
     (SELECT user_id FROM users ORDER BY random() LIMIT 10) AS random_users
LIMIT 10;
