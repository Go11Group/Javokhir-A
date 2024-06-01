-- Active: 1717073257299@@127.0.0.1@5432@leetcode_sql50@public

CREATE TYPE state AS ENUM('approved', 'declined');
Create table If Not Exists Transactions (id int, country varchar(4), state STATE, amount int, trans_date date)

insert into Transactions (id, country, state, amount, trans_date) values ('121', 'US', 'approved', '1000', '2018-12-18');
insert into Transactions (id, country, state, amount, trans_date) values ('122', 'US', 'declined', '2000', '2018-12-19');
insert into Transactions (id, country, state, amount, trans_date) values ('123', 'US', 'approved', '2000', '2019-01-01');
insert into Transactions (id, country, state, amount, trans_date) values ('124', 'DE', 'approved', '2000', '2019-01-07');

SELECT *from transactions;


SELECT to_char(t1.trans_date, 'YYYY-MM') AS month, 
t1.country,
COUNT(*) AS trans_count,
SUM((CASE WHEN t1.state = 'approved' THEN 1 ELSE 0 END)) AS approved_count,
SUM(t1.amount) AS trans_total_amount,
SUM(CASE WHEN t1.state = 'approved' THEN t1.amount ELSE 0 END) AS approved_total_amount
FROM transactions t1
JOIN transactions t2 ON EXTRACT(YEAR FROM t1.trans_date) = EXTRACT(YEAR FROM t2.trans_date)
                     AND EXTRACT(MONTH FROM t1.trans_date) = EXTRACT(MONTH FROM t2.trans_date)
                     AND t1.id = t2.id
GROUP BY month, t1.country;
-- ORDER BY t1.trans_date;
