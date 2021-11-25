CREATE TABLE IF NOT EXISTS transactions (
    transaction_id UUID PRIMARY KEY,
    currency_name VARCHAR(64),
    amount INTEGER,
    price   FLOAT,
    transaction_time   VARCHAR(64)
);

create or replace function insert_notify()
 returns trigger
 language plpgsql
as $$
declare
channel text := TG_ARGV[0];
begin
  PERFORM (
     with resp(transaction_id, currency_name, amount, price, transaction_time) as
     (
       select NEW.transaction_id, NEW.currency_name, NEW.amount, NEW.price, NEW.transaction_time
     )
     select pg_notify(channel, row_to_json(resp)::text)
       from resp
  );
RETURN NULL;
end;
$$;

CREATE TRIGGER insert_transaction
    AFTER INSERT
    ON transactions
    FOR EACH ROW
    EXECUTE PROCEDURE insert_notify('transactions');


-- create or replace function delete_notify()
--     returns trigger
--     language plpgsql
-- as $$
-- declare
--     channel text := TG_ARGV[0];
-- begin
--     PERFORM (
--         with resp(transaction_id, currency_name, amount, price, transaction_time) as
--                  (
--                      select NEW.transaction_id, NEW.currency_name, NEW.amount, NEW.price, NEW.transaction_time
--                  )
--         select pg_notify(channel, row_to_json(resp)::text)
--         from resp
--     );
--     RETURN NULL;
-- end;
-- $$;
--
-- CREATE TRIGGER delete_transaction
--     AFTER DELETE
--     ON transactions
--     FOR EACH ROW
--     EXECUTE PROCEDURE delete_notify('transactions');
