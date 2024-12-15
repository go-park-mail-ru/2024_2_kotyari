ALTER TABLE products
    DROP COLUMN if exists "tags";

ALTER TABLE products
    DROP COLUMN if exists "type";
