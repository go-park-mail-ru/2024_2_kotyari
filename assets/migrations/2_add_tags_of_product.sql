ALTER TABLE products
    ADD COLUMN "tags" text[] DEFAULT '{}';