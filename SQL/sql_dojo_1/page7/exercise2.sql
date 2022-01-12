-- dapatkan total penjualan dan total laba untuk seluruh site
SELECT SUM(price) AS "total penjualan", SUM(price - cost) AS "total laba"
FROM items
JOIN sales_records
ON items.id = sales_records.item_id;