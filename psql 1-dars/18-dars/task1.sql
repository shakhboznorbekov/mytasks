
            --  BANKDA PUL O`TKAZISH
create table users1 (
    user_id serial primary key,
    name varchar(24),
    balance decimal
);
create table bank_shoot (
    bank_id serial primary key,
    bank_balance decimal 
);

insert into bank_shoot(bank_balance)values(0);
insert into users1(name,balance) values
('Shaxboz', 20000),
('Jamshid', 30000),
('Sarvar', 3000);


create or replace procedure payment(sender int, receiver int, price decimal) language plpgsql
as

$$
    begin
        if (select 
                balance
            from 
                users1
                where users1.user_id =sender) >= price then 
            update users1 set balance= balance-price-price*0.01 where user_id =sender;
            update bank_shoot set bank_balance= bank_balance+price*0.01 where bank_id=1;
            update users1 set balance= balance+price where user_id =receiver;
            commit;
        
        else 
            raise info 'not enought money';
        rollback;
        end if;
    end;
$$;






























-- $$
--     begin
--         if (select 
--                 balance
--             from 
--                 users1
--                 where users1.user_id =sender) >= price then 
--             update users1 set balance= balance-price where user_id =sender;
--             update users1 set balance= balance+price where user_id =receiver;
--             commit;
--             return;
--         end if; 
--         raise info 'not enought money';
--         rollback;
--     end;
-- $$;






-- $$
--     begin
--         if (select 
--                 balance
--             from 
--                 users1
--                 where users1.user_id =sender) >= price then 
--             update users1 set balance= balance-price-price*0.01 where user_id =sender;
--             update bank_shoot set bank_balance= bank_balance+price*0.001 where bank_id =receiver;
--             commit;
        
--         else 
--             raise info 'not enought money';
--         rollback;
--         end if;
--     end;
-- $$;