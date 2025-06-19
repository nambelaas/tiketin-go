-- +migrate Up
create table
    tickets (
        id serial primary key,
        event_id integer not null,
        name varchar,
        price numeric(10, 2),
        quota integer,
        created_at timestamp default CURRENT_TIMESTAMP,
        modified_at timestamp,
        foreign key (event_id) references events(id) on delete cascade
    );

-- +migrate Down
drop table tickets;