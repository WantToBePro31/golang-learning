-- TODO: answer here
SELECT o.id AS order_id, fullname, email, product_name, unit_price, quantity, order_date
FROM orders o JOIN users u on o.user_id = u.id
WHERE status = 'active' and (unit_price*quantity > 500000 or quantity > 20);
