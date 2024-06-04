-- Active: 1717073257299@@127.0.0.1@5432@leetcode_sql50@public
Create table If Not Exists Teacher (teacher_id int, subject_id int, dept_id int);
insert into Teacher (teacher_id, subject_id, dept_id) values 
 ('1', '2', '3'),
('1', '2', '4'),
('1', '3', '3'),
('2', '1', '1'),
('2', '2', '1'),
('2', '3', '1'),
('2', '4', '1');

TRUNCATE TABLE teacher;
SELECT *FROM teacher;


SELECT teacher_id, COUNT(DISTINCT subject_id) AS cnt
FROM teacher  
GROUP BY teacher_id;
