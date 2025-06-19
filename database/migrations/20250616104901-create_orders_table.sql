-- +migrate Up
create type order_status as ENUM ('new', 'paid', 'complete', 'cancelled');

create table
    orders (
        id serial primary key,
        user_id integer not null,
        event_id integer not null,
        status order_status default 'new',
        total_price numeric(10, 2),
        paid_at timestamp,
        payment_method varchar,
        created_at timestamp default CURRENT_TIMESTAMP,
        modified_at timestamp,
        foreign key (user_id) references users (id) on delete cascade,
        foreign key (event_id) references events (id) on delete cascade
    );

-- +migrate Down
drop table order;

drop type order_status;