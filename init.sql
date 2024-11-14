CREATE TABLE users (
    id bigserial primary key,
    name varchar(255),
    email varchar(255) unique,
    password varchar(255),
    role varchar(255)
);

CREATE TABLE torrents (
    id bigserial primary key,
    name varchar(255),
    status varchar(255),
    filepath varchar(255),
    output_directory varchar(255),
    created timestamp,
    updated timestamp
);