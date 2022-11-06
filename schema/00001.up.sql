create table schools (
    id text primary key,
    name text,
    phone_number text,
    email text unique,
    password_hash text
);

create table classes (
    id serial primary key,
    school_id text references schools(id),
    number integer,
    letter text
);

create table students (
    id serial primary key,
    first_name text,
    last_name text,
    middle_name text,
    class_id integer
        references classes(id),
    school_id text
        references schools(id),
    email text unique,
    phone_number text unique
);

create index on schools(id);
create index on schools(email);

create unique index on classes (school_id, number, letter)