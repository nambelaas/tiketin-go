-- +migrate Up
create type order_status as ENUM ('new', 'complete', 'cancelled');
create table
    orders (
        id serial primary key,
        order_id varchar,
        user_id integer not null,
        ticket_id integer not null,
        status order_status default 'new',
        quantity int,
        total numeric(10, 2),
        paid_at timestamp,
        payment_method varchar,
        is_check_in boolean,
        checked_in_at timestamp,
        created_at timestamp default CURRENT_TIMESTAMP,
        foreign key (user_id) references users (id) on delete cascade,
        foreign key (ticket_id) references tickets (id) on delete cascade
    );

-- +migrate Down
drop table order;

drop type order_status;