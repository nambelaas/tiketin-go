-- +migrate Up
create type user_role as ENUM ('user', 'organizer', 'admin');

create table
    users (
        "id" serial PRIMARY KEY,
        "name" varchar,
        "email" varchar UNIQUE,
        "password_hash" varchar,
        "role" user_role default 'user',
        "created_at" timestamp
    );

-- +migrate Down
drop table users;

drop type user_role;