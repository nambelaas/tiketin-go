-- +migrate Up
create type user_role as ENUM ('user', 'organizer', 'admin');

create table
    users (
        "id" serial PRIMARY KEY,
        "name" varchar,
        "email" varchar UNIQUE,
        "password" varchar,
        "role" user_role default 'user',
        "created_at" timestamp default CURRENT_TIMESTAMP
    );

-- +migrate Down
drop table users;

drop type user_role;