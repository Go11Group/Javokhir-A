CREATE TABLE books (
    book_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    author VARCHAR NOT NULL,
    total_copies INT NOT NULL
);

CREATE TABLE Members(
    member_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    join_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE Loans(
    loan_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID REFERENCES Books(book_id),
    memeber_id UUID REFERENCES Memeber(member_id),
    loan_date DATE NOT NULL DEFAULT gen_random_uuid(),
    return_date DATE NOT NULL
);
INSERT INTO books (title, author, total_copies) VALUES
('The Great Gatsby', 'F. Scott Fitzgerald', 5),
('To Kill a Mockingbird', 'Harper Lee', 3),
('1984', 'George Orwell', 7),
('Pride and Prejudice', 'Jane Austen', 4),
('The Catcher in the Rye', 'J.D. Salinger', 6),
('Lord of the Flies', 'William Golding', 2),
('The Hobbit', 'J.R.R. Tolkien', 8),
('Animal Farm', 'George Orwell', 4),
('The Grapes of Wrath', 'John Steinbeck', 3),
('Brave New World', 'Aldous Huxley', 5);

INSERT INTO Members (name, email) VALUES
('Alice Smith', 'alice@example.com'),
('Bob Johnson', 'bob@example.com'),
('Charlie Brown', 'charlie@example.com'),
('David Lee', 'david@example.com'),
('Emily Davis', 'emily@example.com'),
('Frank Wilson', 'frank@example.com'),
('Grace Taylor', 'grace@example.com'),
('Henry Clark', 'henry@example.com'),
('Ivy Anderson', 'ivy@example.com'),
('Jack Evans', 'jack@example.com');


SELECT m.name AS member_name, b.title AS book_title
FROM Members m
LEFT JOIN Loans l ON m.member_id = l.member_id
LEFT JOIN Books b ON l.book_id = b.book_id;


UPDATE  Loans SET return_date = CURRENT_DATE + INTERVAL '30 days';