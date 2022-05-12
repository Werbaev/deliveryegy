create unique index users_login_uindex
    on users (login);

create unique index couriers_login_uindex
    on couriers (login);

create unique index vendor_users_login_uindex
    on vendor_users (login);
