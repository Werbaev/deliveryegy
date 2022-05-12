alter table if exists categories add column if not exists merchant_id integer constraint categories_merchant_id_fk references merchants;
update categories set merchant_id = (select id from merchants order by merchants.created_at limit 1);