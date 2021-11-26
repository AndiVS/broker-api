CREATE TABLE IF NOT EXISTS positions (
    positions_id UUID PRIMARY KEY,
    currency_name VARCHAR(64),
    amount INTEGER,
    open_price   FLOAT,
    open_time   VARCHAR(64),
    close_price   FLOAT,
    close_time   VARCHAR(64)
);

create or replace function insert_notify()
 returns trigger
 language plpgsql
as $$
declare
channel text := TG_ARGV[0];
begin
  PERFORM (
     with resp(positions_id, currency_name, amount, open_price, open_time) as
     (
       select NEW.positions_id, NEW.currency_name, NEW.amount, NEW.open_price, NEW.open_time
     )
     select pg_notify(channel, row_to_json(resp)::text)
       from resp
  );
RETURN NULL;
end;
$$;

CREATE TRIGGER open_position
    AFTER INSERT
    ON positions
    FOR EACH ROW
    EXECUTE PROCEDURE insert_notify('positions');


/*create or replace function update_notify()
    returns trigger
    language plpgsql
as $$
declare
    channel text := TG_ARGV[0];
begin
    PERFORM (
        with resp(positions_id, currency_name, amount, open_price, open_time) as
                 (
                     select NEW.positions_id, NEW.currency_name, NEW.amount, NEW.open_price, NEW.open_time
                 )
        select pg_notify(channel, row_to_json(resp)::text)
        from resp
    );
    RETURN NULL;
end;
$$;

CREATE TRIGGER close_position
    AFTER DELETE
    ON positions
    FOR EACH ROW
    EXECUTE PROCEDURE update_notify('positions');
*/