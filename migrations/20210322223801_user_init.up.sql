CREATE EXTENSION if not exists pgcrypto;

create table "user"
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name text default '',
    email text default '',
    password bytea not null,
    salt bytea not null,
    phone text default ''
);
