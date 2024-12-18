CREATE TABLE wish_list_links (
    id bigint generated always as identity,
    link UUID not null ,
    user_id BIGINT NOT NULL,

    primary key (id),
    foreign key (user_id) references "users"("id")
);
