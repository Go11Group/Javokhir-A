psql -U postgres -d postgres -h localhost // entering p sql

CREATE DATABASE LESSON23; // Creating new database.

\c lesson23; // Connecting to the new database.

CREATE TABLE CARS(
	ID SERIAL PRIMARY KEY,
	BRAND VARCAHR NOT NULL,
	MODEL VARCHAR NOT NULL,
	CREATED_YEAR INT,
	PRICE FLOAT
); // Creating new table called cars.

ALTER TABLE CARS 
	ALTER COLUMN CREATED_YEAR SET NOT NULL; // Altering COLUMN TO SET NOT NULL.
	
// inserting data for cars table
INSERT INTO CARS (BRAND, MODEL, CREATED_YEAR, PRICE) VALUES
	('Toyota', 'Camry', 2010, 15000),
	('Honda', 'Accord', 2012, 16000),
	('Ford', 'Mustang', 2015, 25000),
	('Chevrolet', 'Malibu', 2013, 14000),
	('Nissan', 'Altima', 2014, 14500),
	('BMW', '3 Series', 2018, 30000),
	('Audi', 'A4', 2017, 28000),
	('Mercedes-Benz', 'C-Class', 2019, 35000),
	('Hyundai', 'Elantra', 2016, 13000),
	('Kia', 'Optima', 2020, 20000);


// ORDER BY is used to sort records

SELECT *FROM CARS 
	ORDER BY ID DESC; //this query sellects all records and sorts in desending order

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    age INT
);

INSERT INTO users (first_name, last_name, age) VALUES
('John', 'Doe', 25),
('Jane', 'Smith', 30),
('Michael', 'Johnson', 35),
('Emily', 'Davis', 28),
('Chris', 'Brown', 22),
('Amanda', 'Wilson', 26),
('David', 'Taylor', 32),
('Sarah', 'Moore', 27),
('James', 'Anderson', 29),
('Patricia', 'Thomas', 31);

SELECT u.id, u.first_name, u.last_name, u.age, 
		c.Brand, c.model, c.price, c.created_year
FROM Users u
	JOIN Cars c ON c.user_id = u.id 
	WHERE c.user_id = NULL;


SELECT first_name 
	FROM users GROUP BY first_name
		HAVING LENGTH(first_name) = (
			SELECT MAX(LENGTH(first_name))
				FROM users
		);