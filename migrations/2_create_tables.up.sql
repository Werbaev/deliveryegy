CREATE TABLE IF NOT EXISTS merchants
(
    id                  serial not null
                        constraint merchants_pk
                        primary key,
    guid                uuid      default uuid_generate_v4(),
    name                varchar,
    logo                varchar,
    background_image    varchar,
    comission           numeric,
    status              boolean,
    delivery_time       numeric,
    created_at          timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS merchant_branches
(
    id                  serial not null
                        constraint branches_pk
                        primary key,
    guid                uuid      default uuid_generate_v4(),
    name                varchar,
    address             varchar,
    merchant_id         integer not null constraint branches_merchants_id_fk
                        references merchants,
    created_at          timestamp default CURRENT_TIMESTAMP
);