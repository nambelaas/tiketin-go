-- +migrate Up
create table
    event_types (
        id serial primary key,
        name varchar not null,
        created_at timestamp default CURRENT_TIMESTAMP,
        modified_at timestamp
    );

-- +migrate Down
drop table event_types;