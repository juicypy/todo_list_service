create table task_comments
(
    id uuid unique not null default gen_random_uuid(),
    user_id uuid not null,
    task_id uuid not null,
    comment text default '' not null,
    created_at integer
);