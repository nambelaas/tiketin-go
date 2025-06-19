-- +migrate Up
create table
    events (
        id serial primary key,
        user_id integer not null,
        title varchar,
        description varchar,
        location varchar,
        event_date timestamp,
        event_type_id integer not null,
        status varchar,
        created_at timestamp default CURRENT_TIMESTAMP,
        modified_at timestamp,
        foreign key (user_id) references users (id) on delete cascade,
        foreign key (event_type_id) references event_types (id) on delete cascade
    );

-- +migrate Down
drop table events;