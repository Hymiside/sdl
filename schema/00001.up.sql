create table schools (
    id text primary key,
    name text unique,
    phone_number text,
    email text,
    password_hash text
);

create index on schools(id);
create index on schools(name);