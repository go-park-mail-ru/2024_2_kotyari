CREATE USER service_account WITH PASSWORD 'kotyari-Password123@';

GRANT CONNECT ON DATABASE kotyari_2024 TO service_account;

CREATE ROLE service_account_role WITH LOGIN;

GRANT USAGE ON SCHEMA public TO service_account_role;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO service_account_role;
GRANT INSERT ON orders, product_orders, carts TO service_account_role;
GRANT UPDATE ON users, carts, products TO service_account_role;

GRANT service_account_role TO service_account;