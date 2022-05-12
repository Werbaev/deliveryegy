CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id          serial not null
        constraint users_pk
            primary key,
    guid        uuid default uuid_generate_v4(),
    name        varchar,
    login       varchar,
    password    varchar,
    created_at  timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories
(
    id         serial not null
        constraint categories_pk
            primary key,
    guid       uuid      default uuid_generate_v4(),
    created_at timestamp default CURRENT_TIMESTAMP,
    name        varchar
);

CREATE TABLE IF NOT EXISTS products
(
    id         serial not null
        constraint products_pk
            primary key,
    guid       uuid      default uuid_generate_v4(),
    created_at timestamp default CURRENT_TIMESTAMP,
    name        varchar,
    category_id integer not null constraint products_category_id_fk
            references categories,
    price       numeric
);

CREATE TABLE IF NOT EXISTS orders
(
    id         serial not null
        constraint orders_pk
            primary key,
    guid       uuid      default uuid_generate_v4(),
    created_at timestamp default CURRENT_TIMESTAMP,
    user_id integer not null constraint orders_users_id_fk
            references users,
    product_id integer not null constraint orders_products_id_fk
            references products,
    comment        varchar,
    address        varchar
);