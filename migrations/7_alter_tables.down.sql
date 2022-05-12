alter table if exists orders drop column if exists courier_id;
alter table if exists orders drop column if exists branch_id;
alter table if exists products drop column if exists image;

alter table if exists orders drop column if exists products;

alter table orders
    alter column product_id set not null;

alter table if exists orders drop column if exists total_price;

alter table if exists orders
rename column delivery_price to total_price;

alter table if exists orders drop column if exists payment_type;

alter table if exists orders drop column if exists delivery_type;

alter table if exists users drop column if exists phone_number;

alter table if exists merchants drop column if exists description;