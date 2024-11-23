create type survey_type as enum (
    'site',
    'purchase'
    );

create table if not exists surveys (
    id bigint generated always as identity,
    user_id bigint not null,
    text text,
    rating smallint,
    type survey_type,
    updated_at timestamp with time zone not null default current_timestamp,
    created_at timestamp with time zone not null default current_timestamp,
    primary key (id)
)