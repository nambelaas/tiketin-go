-- +migrate Up
create table
    order_items (
        id serial primary key,
        order_id int not null,
        ticket_type_id int not null,
        quantity int,
        qr_code_url varchar,
        is_check_in boolean default false,
        created_at timestamp default CURRENT_TIMESTAMP,
        modified_at timestamp,
        foreign key (order_id) references orders (id) on delete cascade,
        foreign key (ticket_type_id) references tickets (id) on delete cascade
    );

-- +migrate Down
drop table order_items;