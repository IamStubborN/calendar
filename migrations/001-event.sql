set search_path=public;
create schema if not exists public;
create table if not exists events
(
    id bigserial not null
        constraint events_pk
            primary key,
    name varchar(50),
    description text,
    date date not null
);

create unique index if not exists events_id_uindex
    on events (id);