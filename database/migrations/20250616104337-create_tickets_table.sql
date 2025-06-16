-- +migrate Up
create table
    tickets (
        id serial primary key,
        tiket_id varchar,
        event_id integer not null,
        name varchar,
        price numeric(10, 2),
        quota integer,
        created_at timestamp default CURRENT_TIMESTAMP,
        foreign key (event_id) references events(id) on delete cascade
    );

-- +migrate Down
drop table tickets;