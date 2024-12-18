ALTER TABLE products
    ADD COLUMN "tags" text[] DEFAULT '{}';

ALTER TABLE products
    ADD COLUMN "type" text NOT NULL default '';
