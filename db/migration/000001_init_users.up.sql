create schema if not exists users authorization administrator;

create table if not exists users.accounts
(
    id       serial not null unique,
    nickname text   not null unique
)