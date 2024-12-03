create table if not exists user_promocodes(
    id bigint generated always as identity,
    user_id bigint not null,
    promo_id bigint not null,
    updated_at timestamp with time zone not null default current_timestamp,
    created_at timestamp with time zone not null default current_timestamp,
    primary key (id),
    foreign key (user_id) references "users"(id),
    foreign key (promo_id) references "promocodes"(id)
);

create table if not exists promocodes(
    id bigint generated always as identity,
    name text not null,
    bonus smallint not null,
    updated_at timestamp with time zone not null default current_timestamp,
    created_at timestamp with time zone not null default current_timestamp,
    primary key (id)
);


create index if not exists idx_promocodes_user_id on user_promocodes(user_id);
create index if not exists idx_promocodes_promo_id on user_promocodes(promo_id);
