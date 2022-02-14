ERD yang disediakan akan lebih baik jika di table outlet ditambahkan user_id sehingga dapat lebih efisien dalam query dengan menambahkan kondisi di joins outlet

SELECT t.id, t.bill_total ,m.merchant_name, o.outlet_name, s.date  
FROM (SELECT i::date AS date FROM generate_series('2021-11-01 00:00:00', '2021-11-30 00:00:00', interval '1 day') i) s  
LEFT JOIN transactions t on t.created_at::date = s.date
LEFT JOIN merchants m on t.merchant_id = m.id and m.user_id = 1 
LEFT JOIN outlets o on o.merchant_id = m.id and m.user_id = 1 
WHERE s.date >= '2021-11-01 00:00:00' and s.date <= '2021-11-30 00:00:00' 
GROUP BY t.id,m.merchant_name, o.outlet_name, s.date 
ORDER BY s.date asc LIMIT 100