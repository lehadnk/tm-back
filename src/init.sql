CREATE TABLE users (
    id bigserial primary key,
    name varchar(255),
    email varchar(255) unique,
    password varchar(255),
    role varchar(255)
);