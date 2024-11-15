create table if not exists reviews (
    id bigint generated always as identity,
    product_id bigint not null,
    user_id bigint not null,
    text text not null,
    good_review_count int default 0,
    bad_review_count int default 0,
    updated_at timestamp with time zone not null default current_timestamp,
    created_at timestamp with time zone not null default current_timestamp,
    primary key (id),
    foreign key (product_id) references "products"("id"),
    foreign key (user_id) references "users"("id")
)