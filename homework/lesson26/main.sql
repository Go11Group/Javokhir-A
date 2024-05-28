create table department
(
    id   uuid primary key default gen_random_uuid() not null,
    name varchar                                    not null
);


create table if not exists employee
(
    id            uuid primary key default gen_random_uuid() not null,
    name          varchar not null,
    department_id uuid references department (id),
    salary        int     not null default 0
);


create table if not exists task
(
    id          uuid primary key default gen_random_uuid() not null,
    number      int unique                                 not null,
    description text             default ''
);

create table if not exists employee_task
(
    id          uuid primary key default gen_random_uuid() not null,
    employee_id uuid references employee (id),
    task_id     uuid references task (id)
);

insert into department(id, name) values('5939ed2c-4b4b-43ae-a0ab-695175c3a731','HR');
insert into department(id, name) values('558f6658-7793-4107-8930-cbf5dbd2b008','Programming');
insert into department(id, name) values('08894581-f43e-4713-bccd-4283b53fd08b','Marketing');

INSERT INTO employee (name, department_id, salary) VALUES
   ('John Doe', '5939ed2c-4b4b-43ae-a0ab-695175c3a731', 50000),
   ('Alice Smith', '558f6658-7793-4107-8930-cbf5dbd2b008', 60000),
   ('Bob Johnson', '558f6658-7793-4107-8930-cbf5dbd2b008', 55000),
   ('Emily Davis', '08894581-f43e-4713-bccd-4283b53fd08b', 52000),
   ('Michael Brown', '08894581-f43e-4713-bccd-4283b53fd08b', 62000),
   ('Jennifer Lee', '08894581-f43e-4713-bccd-4283b53fd08b', 58000),
   ('David Wilson', '5939ed2c-4b4b-43ae-a0ab-695175c3a731', 54000),
   ('Jessica Martinez', '5939ed2c-4b4b-43ae-a0ab-695175c3a731', 67000);

INSERT INTO task (id, number, description) VALUES
   ('1d7ab001-df4e-4ac6-adf2-70cab7568c42', 101, 'Task 101 description'),
   ('fbb3cd9f-0567-4496-b694-96de4aad5c8d', 102, 'Task 102 description'),
   ('488fb943-37c3-47bb-b3af-1cdd5b9dbfbb', 103, 'Task 103 description');

INSERT INTO employee_task (employee_id, task_id) VALUES
     ('0a876d4f-937e-4f77-9383-825f260728b2','1d7ab001-df4e-4ac6-adf2-70cab7568c42'),
     ('45c2aabd-143d-40a5-b053-2142ab6521f9','1d7ab001-df4e-4ac6-adf2-70cab7568c42'),
     ('45c2aabd-143d-40a5-b053-2142ab6521f9','fbb3cd9f-0567-4496-b694-96de4aad5c8d'),
     ('0a876d4f-937e-4f77-9383-825f260728b2','488fb943-37c3-47bb-b3af-1cdd5b9dbfbb'),
     ('45c2aabd-143d-40a5-b053-2142ab6521f9','488fb943-37c3-47bb-b3af-1cdd5b9dbfbb');

select array_agg(e.name), t.description from employee e
 left join department d on e.department_id=d.id
 left join employee_task et on e.id = et.employee_id
 left join task t on t.id = et.task_id
where t.description is not null
group by t.description;

select array_agg(e.name), d.name, avg(salary) sum from employee e
    left join department d on d.id = e.department_id
    group by d.id
    order by sum
;

-- Student table
-- course table
-- student_course table
-- grade table

-- 1. grade tabledagi student_id va course_id o'rniga student_course tabledagi
-- idni olishimiz darkor;

-- 2. guruhdagi eng yaxshi o'qiydigan studentlarni har guruh bo'yicha chiqaring. Agarda
-- eng yaxshi baholar bir nechta kishida bo'lsa, hammasi chiqsin.

-- 3. guruhning o'rtacha bahosini har bir guruh bo'yicha chiqaring;

-- 4. eng yosh o'quvchi guruh bo'yicha chiqarilsin chiqarilsin;
-- yani gar bir guruhdan eng yosh o'quvchi(bir nechta bo'lsa har birini)

-- 5. eng yaxshi o'qiydigan guruh chiqarilsin
-- ya'ni, har bir guruhning o'rtacha bahosining eng katta(yaxshi) bo'lgani


create table courses 
(   
    course_id uuid primary key default gen_random_uuid(),
    course_name varchar not null
);

create table students
(
    student_id uuid primary key default gen_random_uuid(),
    student_name varchar not null
);

create table student_course
(
    id uuid primary key default gen_random_uuid(),
    student_id uuid references studens(student_id),
    course_id uuid references courses(course_id)
);


create table grades
(
    grade_id uuid primary key default gen_random_uuid(),
    student_course_id uuid references student_course(id),
    grade int check (grade <= 100 and grade >= 1)
);

INSERT INTO courses (course_name)
VALUES ('Introduction to Programming'),
       ('Data Analysis and Visualization'),             
       ('Web Development Fundamentals'),
       ('Machine Learning for Beginners'),
       ('Software Engineering Principles'),
       ('Creative Writing Workshop'),
       ('Public Speaking and Communication'),
       ('Calculus I: Differentiation and Integration'),
       ('Physics for Scientists and Engineers'),
       ('World History: From Antiquity to the Present');
3
INSERT INTO students (student_name)
VALUES ('Alice Johnson'),
       ('Bob Williams'),
       ('Charlie Brown'),
       ('David Miller'),
       ('Emily Garcia'),
       ('Frederick Anderson'),
       ('Grace Lee'),
       ('Henry Moore'),
       ('Isabella Davis'),
       ('Jacob Wilson');

INSERT INTO student_course(Student_id, course_id) VALUES
('62e30520-26e3-405f-b8b2-386e3ff071c2', '2b4d27e9-4b49-421c-8eb2-ed98a1094ee0'),
('62e30520-26e3-405f-b8b2-386e3ff071c2', '86dc19da-494e-4f9d-af41-21d8c02395c5'),
('6496ce48-2260-48a2-9af9-56fbeb8d070c', 'fa2099e2-b67f-4eab-9f24-fccad64f5804'),
('62e30520-26e3-405f-b8b2-386e3ff071c2', '06d23597-add6-4ff5-950d-b64c1efb76b3'),
('91700992-8d1d-4559-a791-d141d1619f29', '86edb53b-81ab-42ee-8bf7-d5c3b031d028'),
('91700992-8d1d-4559-a791-d141d1619f29', '86dc19da-494e-4f9d-af41-21d8c02395c5'),
('753f167f-4b3a-44f0-ac1c-6bdd8f15fc17', '86dc19da-494e-4f9d-af41-21d8c02395c5'),
('753f167f-4b3a-44f0-ac1c-6bdd8f15fc17', '06d23597-add6-4ff5-950d-b64c1efb76b3'),
('753f167f-4b3a-44f0-ac1c-6bdd8f15fc17', '208734ae-f4fd-45b1-9799-94b831ba7ce9'),
('753f167f-4b3a-44f0-ac1c-6bdd8f15fc17', '86edb53b-81ab-42ee-8bf7-d5c3b031d028'),
('015677ca-80d7-444a-8c9a-db058b623d69', '6e1168b1-8724-44b1-9659-02a8bcecbaed'),
('015677ca-80d7-444a-8c9a-db058b623d69', '0bac3f1d-1f7a-4d7d-b759-fda6c3e80b51'),
('015677ca-80d7-444a-8c9a-db058b623d69', '06d23597-add6-4ff5-950d-b64c1efb76b3'),
('8dd6afc7-04d3-4627-850c-af6b072e4623', '208734ae-f4fd-45b1-9799-94b831ba7ce9'),
('7eb0d90d-d29e-4e7f-b236-4d5c1643187c', 'fa2099e2-b67f-4eab-9f24-fccad64f5804'),
('7b321c32-79b3-42a7-916c-257e466b8737', '0bac3f1d-1f7a-4d7d-b759-fda6c3e80b51'),
('71e0fcbc-f94b-4423-b71e-dde6c314b9c9', '06d23597-add6-4ff5-950d-b64c1efb76b3'),
('38e4e8ed-f494-40d5-9484-57dc60cdcb54', '86dc19da-494e-4f9d-af41-21d8c02395c5');

    
INSERT INTO grades (student_course_id, grade) VALUES
('d28d5fe8-2f9b-430a-927e-f8c46ae5aaad', 80),
('91fd44f1-d9aa-498d-b881-9ae39f3b2ca2', 85),
('0e52f2be-754c-40f3-a551-3cf9ab41effa', 71),
('9625f5fd-2289-4d20-99c7-cd3954e298b5', 98),
('1b7e1770-124d-4eb0-ad62-4ac8189e555a', 65),
('f521e597-18a4-489c-ab2c-dfa3e16855dc', 74),
('5757f008-0705-4109-885c-296c3377c60f', 56),
('7e4d5104-a61e-4ae4-9978-23f4a7bd9659', 98),
('c0d22dbd-05a3-48ca-9383-15b94a207d98', 77),
('4a1d7e6e-d798-4474-84b1-c66bdfcb012c', 73),
('d8705651-23d1-4335-bac9-bb73a7ec3b26', 85),
('3b3fb73c-509c-425c-a1fc-b3d19822df7e', 47),
('2f74d48d-6d65-4d4c-9176-efbcc59bb386', 90),
('a4734ffa-a13b-4fd4-b4b3-bd9a7f431caa', 78),
('3181acbb-6b50-4f70-a41b-5c00dac6966b', 75),
('dc4954ff-af05-4444-a75e-a75f23a6ecef', 48),
('2e4d02da-181e-4f09-849d-03420639f52a', 69),
('38e75358-adfd-4286-a305-92f8ac1a1840', 58);


SELECT s.student_name, c.course_name, g.grade
FROM students s
LEFT JOIN student_course cs
ON s.student_id = cs.student_id
LEFT JOIN courses c 
ON cs.course_id = c.course_id
LEFT JOIN grades g
ON g.student_course_id = cs.id
ORDER BY g.grade DESC;



-- TASK 3
-- 3. guruhning o'rtacha bahosini har bir guruh bo'yicha chiqaring;

SELECT c.course_name, ROUND(COALESCE(AVG(g.grade), 0)::NUMERIC, 2) AS average_grade
FROM courses c
LEFT JOIN student_course cs
ON cs.course_id = c.course_id
LEFT JOIN students s
ON s.student_id = cs.student_id
LEFT JOIN grades g
ON g.student_course_id = cs.id
GROUP BY c.course_name
ORDER BY average_grade DESC;

-- RESULT                  
--------------------------------------------------------------------                   |
--                  course_name                  | average_grade    |                  |
-- ----------------------------------------------+---------------   |                  |
--  Public Speaking and Communication            |         88.75    |                  |
--  Data Analysis and Visualization              |         85.00    |                  |
--  World History: From Antiquity to the Present |         80.00    |                  |
--  Physics for Scientists and Engineers         |         77.50    |                  |
--  Web Development Fundamentals                 |         73.00    |                  |
--  Software Engineering Principles              |         69.00    |                  |
--  Introduction to Programming                  |         68.25    |                  |
--  Machine Learning for Beginners               |         47.50    |                  |
--  Creative Writing Workshop                    |          0.00    |                  |
--  Calculus I: Differentiation and Integration  |          0.00    |                  |
--------------------------------------------------------------------                   |
                    
----------------------------------------------------------------------------------------

-- TASK 4
-- 4. eng yosh o'quvchi guruh bo'yicha chiqarilsin chiqarilsin;
-- yani gar bir guruhdan eng yosh o'quvchi(bir nechta bo'lsa har birini)

SELECT c.course_name, s.student_name, s.student_age
FROM courses c
JOIN student_course cs ON cs.course_id = c.course_id
JOIN students s ON s.student_id = cs.student_id
WHERE s.student_age = (
    SELECT MIN(s2.student_age)
    FROM students s2
    JOIN student_course cs2 ON s2.student_id = cs2.student_id
    WHERE cs2.course_id = c.course_id
)
ORDER BY c.course_name, s.student_name;

-- RESULT

-----------------------------------------------------------------------------------------
--                  course_name                  |    student_name    | student_age     |
-- ----------------------------------------------+--------------------+-------------    |
--  Data Analysis and Visualization              | Emily Garcia       |          21     |
--  Introduction to Programming                  | Jacob Wilson       |          18     |
--  Machine Learning for Beginners               | Isabella Davis     |          19     |
--  Physics for Scientists and Engineers         | David Miller       |          19     |
--  Physics for Scientists and Engineers         | Frederick Anderson |          19     |
--  Public Speaking and Communication            | Alice Johnson      |          19     |
--  Public Speaking and Communication            | David Miller       |          19     |
--  Software Engineering Principles              | Charlie Brown      |          19     |
--  Software Engineering Principles              | David Miller       |          19     |
--  Web Development Fundamentals                 | Grace Lee          |          19     |
--  World History: From Antiquity to the Present | Alice Johnson      |          19     |
-----------------------------------------------------------------------------------------



-- TASK 5
-- 5. eng yaxshi o'qiydigan guruh chiqarilsin
-- ya'ni, har bir guruhning o'rtacha bahosining eng katta(yaxshi) bo'lgani

SELECT c.course_name, ROUND(avg(g.grade)::NUMERIC, 2) average_grade
FROM courses c
LEFT JOIN student_course cs
ON cs.course_id = c.course_id
LEFT JOIN students s
ON s.student_id = cs.student_id
JOIN grades g
ON g.student_course_id = cs.id
GROUP BY c.course_name
ORDER BY average_grade DESC;

-- RESULT
-------------------------------------------------------------------
--                  course_name                  | average_grade   | 
-- ----------------------------------------------+---------------  |
--  Public Speaking and Communication            |         88.75   |
--  Data Analysis and Visualization              |         85.00   |
--  World History: From Antiquity to the Present |         80.00   |
--  Physics for Scientists and Engineers         |         77.50   |
--  Web Development Fundamentals                 |         73.00   |
--  Software Engineering Principles              |         69.00   |
--  Introduction to Programming                  |         68.25   |
--  Machine Learning for Beginners               |         47.50   |
-------------------------------------------------------------------




-- altering students table to add age column
ALTER TABLE students ADD COLUMN student_age INT DEFAULT 18;

UPDATE students
SET student_age = FLOOR(18 + (RANDOM() * 5))::INT;
-----






--               course_id               |                 course_name                  
-- --------------------------------------+----------------------------------------------
--  86dc19da-494e-4f9d-af41-21d8c02395c5 | Introduction to Programming
--  6e1168b1-8724-44b1-9659-02a8bcecbaed | Data Analysis and Visualization
--  fa2099e2-b67f-4eab-9f24-fccad64f5804 | Web Development Fundamentals
--  0bac3f1d-1f7a-4d7d-b759-fda6c3e80b51 | Machine Learning for Beginners
--  86edb53b-81ab-42ee-8bf7-d5c3b031d028 | Software Engineering Principles
--  dd526c76-772d-4641-b861-bcb11c332e19 | Creative Writing Workshop
--  06d23597-add6-4ff5-950d-b64c1efb76b3 | Public Speaking and Communication
--  5a11344a-e95d-4624-881e-66ca9394c2fa | Calculus I: Differentiation and Integration
--  208734ae-f4fd-45b1-9799-94b831ba7ce9 | Physics for Scientists and Engineers
--  2b4d27e9-4b49-421c-8eb2-ed98a1094ee0 | World History: From Antiquity to the Present


--               student_id              |    student_name    | student_age 
-- --------------------------------------+--------------------+-------------
--  62e30520-26e3-405f-b8b2-386e3ff071c2 | Alice Johnson      |          19
--  6496ce48-2260-48a2-9af9-56fbeb8d070c | Bob Williams       |          22
--  91700992-8d1d-4559-a791-d141d1619f29 | Charlie Brown      |          19
--  753f167f-4b3a-44f0-ac1c-6bdd8f15fc17 | David Miller       |          19
--  015677ca-80d7-444a-8c9a-db058b623d69 | Emily Garcia       |          21
--  8dd6afc7-04d3-4627-850c-af6b072e4623 | Frederick Anderson |          19
--  7eb0d90d-d29e-4e7f-b236-4d5c1643187c | Grace Lee          |          19
--  71e0fcbc-f94b-4423-b71e-dde6c314b9c9 | Henry Moore        |          20
--  7b321c32-79b3-42a7-916c-257e466b8737 | Isabella Davis     |          19
--  38e4e8ed-f494-40d5-9484-57dc60cdcb54 | Jacob Wilson       |          18



--                   id                  |              student_id              |              course_id               
-- --------------------------------------+--------------------------------------+--------------------------------------
--  d28d5fe8-2f9b-430a-927e-f8c46ae5aaad | 62e30520-26e3-405f-b8b2-386e3ff071c2 | 2b4d27e9-4b49-421c-8eb2-ed98a1094ee0
--  91fd44f1-d9aa-498d-b881-9ae39f3b2ca2 | 62e30520-26e3-405f-b8b2-386e3ff071c2 | 86dc19da-494e-4f9d-af41-21d8c02395c5
--  0e52f2be-754c-40f3-a551-3cf9ab41effa | 6496ce48-2260-48a2-9af9-56fbeb8d070c | fa2099e2-b67f-4eab-9f24-fccad64f5804
--  9625f5fd-2289-4d20-99c7-cd3954e298b5 | 62e30520-26e3-405f-b8b2-386e3ff071c2 | 06d23597-add6-4ff5-950d-b64c1efb76b3
--  1b7e1770-124d-4eb0-ad62-4ac8189e555a | 91700992-8d1d-4559-a791-d141d1619f29 | 86edb53b-81ab-42ee-8bf7-d5c3b031d028
--  f521e597-18a4-489c-ab2c-dfa3e16855dc | 91700992-8d1d-4559-a791-d141d1619f29 | 86dc19da-494e-4f9d-af41-21d8c02395c5
--  5757f008-0705-4109-885c-296c3377c60f | 753f167f-4b3a-44f0-ac1c-6bdd8f15fc17 | 86dc19da-494e-4f9d-af41-21d8c02395c5
--  7e4d5104-a61e-4ae4-9978-23f4a7bd9659 | 753f167f-4b3a-44f0-ac1c-6bdd8f15fc17 | 06d23597-add6-4ff5-950d-b64c1efb76b3
--  c0d22dbd-05a3-48ca-9383-15b94a207d98 | 753f167f-4b3a-44f0-ac1c-6bdd8f15fc17 | 208734ae-f4fd-45b1-9799-94b831ba7ce9
--  4a1d7e6e-d798-4474-84b1-c66bdfcb012c | 753f167f-4b3a-44f0-ac1c-6bdd8f15fc17 | 86edb53b-81ab-42ee-8bf7-d5c3b031d028
--  d8705651-23d1-4335-bac9-bb73a7ec3b26 | 015677ca-80d7-444a-8c9a-db058b623d69 | 6e1168b1-8724-44b1-9659-02a8bcecbaed
--  3b3fb73c-509c-425c-a1fc-b3d19822df7e | 015677ca-80d7-444a-8c9a-db058b623d69 | 0bac3f1d-1f7a-4d7d-b759-fda6c3e80b51
--  2f74d48d-6d65-4d4c-9176-efbcc59bb386 | 015677ca-80d7-444a-8c9a-db058b623d69 | 06d23597-add6-4ff5-950d-b64c1efb76b3
--  a4734ffa-a13b-4fd4-b4b3-bd9a7f431caa | 8dd6afc7-04d3-4627-850c-af6b072e4623 | 208734ae-f4fd-45b1-9799-94b831ba7ce9
--  3181acbb-6b50-4f70-a41b-5c00dac6966b | 7eb0d90d-d29e-4e7f-b236-4d5c1643187c | fa2099e2-b67f-4eab-9f24-fccad64f5804
--  dc4954ff-af05-4444-a75e-a75f23a6ecef | 7b321c32-79b3-42a7-916c-257e466b8737 | 0bac3f1d-1f7a-4d7d-b759-fda6c3e80b51
--  2e4d02da-181e-4f09-849d-03420639f52a | 71e0fcbc-f94b-4423-b71e-dde6c314b9c9 | 06d23597-add6-4ff5-950d-b64c1efb76b3
--  38e75358-adfd-4286-a305-92f8ac1a1840 | 38e4e8ed-f494-40d5-9484-57dc60cdcb54 | 86dc19da-494e-4f9d-af41-21d8c02395c5


--     student_name    |                 course_name                  | grade 
-- --------------------+----------------------------------------------+-------
--  Alice Johnson      | Public Speaking and Communication            |    98
--  David Miller       | Public Speaking and Communication            |    98
--  Emily Garcia       | Public Speaking and Communication            |    90
--  Alice Johnson      | Introduction to Programming                  |    85
--  Emily Garcia       | Data Analysis and Visualization              |    85
--  Alice Johnson      | World History: From Antiquity to the Present |    80
--  Frederick Anderson | Physics for Scientists and Engineers         |    78
--  David Miller       | Physics for Scientists and Engineers         |    77
--  Grace Lee          | Web Development Fundamentals                 |    75
--  Charlie Brown      | Introduction to Programming                  |    74
--  David Miller       | Software Engineering Principles              |    73
--  Bob Williams       | Web Development Fundamentals                 |    71
--  Henry Moore        | Public Speaking and Communication            |    69
--  Charlie Brown      | Software Engineering Principles              |    65
--  Jacob Wilson       | Introduction to Programming                  |    58
--  David Miller       | Introduction to Programming                  |    56
--  Isabella Davis     | Machine Learning for Beginners               |    48
--  Emily Garcia       | Machine Learning for Beginners               |    47
-- (18 rows)
