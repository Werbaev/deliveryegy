delete from orders;
alter table if exists orders add column if not exists courier_id integer constraint orders_courier_id_fk references couriers;
alter table if exists orders add column if not exists branch_id integer not null constraint orders_branch_id_fk references merchant_branches;
alter table if exists products add column if not exists image varchar;

alter table if exists orders add column if not exists products jsonb;

alter table orders
    alter column product_id drop not null;

alter table if exists orders add column if not exists total_price numeric default 10000;

alter table if exists orders
rename column total_price to delivery_price;

alter table if exists orders add column if not exists payment_type varchar default 'card';

alter table if exists orders add column if not exists delivery_type varchar default 'delivery';

alter table if exists users add column if not exists phone_number varchar default '998';

alter table if exists merchants add column if not exists description varchar default '.';