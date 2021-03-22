create table task
(
    id uuid unique not null default gen_random_uuid(),
    user_id uuid not null,
    name text default '' not null,
    status integer default 0 not null,
    description text default '' not null,
    date_from integer,
    date_to integer,
    label_ids text[] default '{}',
    modified_at integer,
    created_at integer,
    constraint task_pk
        primary key (id, user_id)
);