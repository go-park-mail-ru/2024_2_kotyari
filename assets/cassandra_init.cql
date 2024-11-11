CREATE KEYSPACE IF NOT EXISTS oxic WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

CREATE TABLE IF NOT EXISTS oxic.carts (
    user_id bigint,
    product_id bigint,
    count int,
    is_selected boolean,
    is_deleted boolean,
    updated_at timestamp,
    created_at timestamp,
    PRIMARY KEY ( user_id, product_id )
);