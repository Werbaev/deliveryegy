CREATE TABLE IF NOT EXISTS couriers
(
    id          serial not null
        constraint couriers_pk
            primary key,
    guid        uuid default uuid_generate_v4(),
    name        varchar,
    login       varchar,
    password    varchar,
    created_at  timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS vendor_users
(
    id                  serial not null
                    constraint vendor_users_pk
                        primary key,
    guid                uuid default uuid_generate_v4(),
    name                varchar,
    login               varchar,
    password            varchar,
    merchant_branch_id  integer not null constraint vendor_users_merchant_branches_id_fk
                        references merchant_branches,
    created_at          timestamp default CURRENT_TIMESTAMP
);

alter table if exists orders add column if not exists status varchar default 'new';