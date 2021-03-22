create table label
(
    id uuid unique not null default gen_random_uuid(),
    name text default '' not null,
    color text default '' not null
);

insert into label (name, color) values ('white', '#FFFFFF');
insert into label (name, color) values ('black', '#000000');
insert into label (name, color) values ('black', '#0000CD');
