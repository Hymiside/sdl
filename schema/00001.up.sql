create table schools (
    id text primary key,
    name text,
    phone_number text,
    email text unique,
    password_hash text
);

create index on schools(id);
create index on schools(email);