-- +migrate Up
create table
    reviews (
        id serial primary key,
        review_id varchar,
        user_id integer not null,
        event_id integer not null,
        rating integer not null,
        comment varchar,
        created_at timestamp default CURRENT_TIMESTAMP,
        modified_at timestamp,
        foreign key (user_id) references users (id) on delete cascade,
        foreign key (event_id) references events (id) on delete cascade
    );

-- +migrate Down
drop table reviews;