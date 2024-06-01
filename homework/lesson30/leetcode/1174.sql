-- Active: 1717073257299@@127.0.0.1@5432@leetcode_sql50@public
Create table If Not Exists Delivery (delivery_id int, customer_id int, order_date date, customer_pref_delivery_date date)

insert into Delivery (delivery_id, customer_id, order_date, customer_pref_delivery_date) values 
('1', '1', '2019-08-01', '2019-08-02'),
('2', '2', '2019-08-02', '2019-08-02'),
('3', '1', '2019-08-11', '2019-08-12'),
('4', '3', '2019-08-24', '2019-08-24'),
('5', '3', '2019-08-21', '2019-08-22'),
('6', '2', '2019-08-11', '2019-08-13'),
('7', '4', '2019-08-09', '2019-08-09')


SELECT *FROM delivery;

SELECT
	ROUND(
		SUM(
			CASE
				WHEN ORDER_DATE = CUSTOMER_PREF_DELIVERY_DATE THEN 1
				ELSE 0
			END
		) * 100.00 / COUNT(*),
		2
	) AS IMMEDIATE_PERCENTAGE
FROM
	DELIVERY
WHERE
	(CUSTOMER_ID, ORDER_DATE) IN (
		SELECT
			CUSTOMER_ID,
			MIN(ORDER_DATE)
		FROM
			DELIVERY
		GROUP BY
    		CUSTOMER_ID
    )
