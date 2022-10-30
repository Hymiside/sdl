create table users (
    id text primary key,
    firstname text,
    lastname text,
    username text unique,
    email text,
    password_hash text
);

create index on users(id);
create index on users(username);