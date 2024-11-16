create table if not exists reviews (
    id bigint generated always as identity,
    product_id bigint not null,
    user_id bigint not null,
    text text,
    rating smallint,
    is_private bool,
    updated_at timestamp with time zone not null default current_timestamp,
    created_at timestamp with time zone not null default current_timestamp,
    primary key (id),
    foreign key (product_id) references "products"("id"),
    foreign key (user_id) references "users"("id")
)