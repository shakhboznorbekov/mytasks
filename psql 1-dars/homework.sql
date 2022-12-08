drop table if exists investor cascade;
drop table if exists modul cascade;

create table investor (
    investor_id serial primary key,
    name varchar(22)
);

create table modul (
    modul_id serial not null primary key,
    modul_name varchar(24)
);

create table add_tariff (
    add_tariff_id serial primary key,
    name varchar(24) not null,
    price numeric not null
);

create table order_payment (
    order_payment_id serial primary key,
    total_price numeric not null,
    add_tariff_price numeric,
    investor numeric not null,
    company numeric not null,
    insurance numeric not null
);




--         ****************
alter table car add column investor_id int references investor(investor_id);
alter table car add column investor_percentage int;
alter table car alter column investor_percentage set not null;

insert into investor (name)
values
('Akbar'),
('Kamoliddin'),
('Farrux');

update car set investor_id = 1 where car_id = 1;
update car set investor_id = 1 where car_id = 2;
update car set investor_id = 2 where car_id = 3;
update car set investor_id = 3 where car_id = 4;
update car set investor_id = 3 where car_id = 5;

insert into modul (modul_name) values
('Malibu'),
('Tracker'),
('Cobalt'),
('Onix'),
('Nexia');

alter table car add column modul_id int references modul(modul_id);
alter table car alter column modul_id set not null;

update car set modul_id = 1 where car_id = 1;
update car set modul_id = 2 where car_id = 2;
update car set modul_id = 3 where car_id = 3;
update car set modul_id = 4 where car_id = 4;
update car set modul_id = 5 where car_id = 5;

alter table car add column insurance numeric;
alter table car alter column insurance set not null;


insert into add_tariff (name, price)
values
('Haydovchi', 300000),
('Anti-Radar', 40000),
('Serena', 30000);

alter table orders add column add_tariff_price_day numeric;
alter table orders add column total_price numeric;
alter table orders alter column total_price set not null;
alter table orders add column is_paid integer default 0;


insert into orders (
    price,
    status,
    car_id,
    client_id,
    day_count,
    add_tariff_price_day,
    total_price
) values 
(
    (
        select 
        day_price + 340000 + 30000 
        from car where car_id = 1
    ),
    'new',
    1,
    1,
    4,
    340000,
    (
        select (day_price + 340000 + 30000) * 4 
        from car where car_id = 1
    )
);


update orders set is_paid = 0 where order_id = 5;
update orders set is_paid = 1 where order_id = 5;





alter table order_payment add column investor_id int not null references investor(investor_id); 


create or replace function investor_share(investor_id int) returns table(
    investor_name varchar(255),
    total_price numeric
) language plpgsql
as
$$
    begin
        return QUERY
        
        select
            i.name,
            op.total_price 
        from order_payment op
        join investor AS i on i.investor_id=order_payment.investor_id;

        insert into orders (
            price,
            status,
            car_id,
            client_id,
            day_count,
            add_tariff_price_day,
            total_price
        ) values 
        (
            (
                select 
                day_price + 340000 + 30000 
                from car where car_id = 1
            ),
            'client_returned ',
            1,
            1,
            4,
            340000,
            (
                select (day_price + 340000 + 30000) * 4 
                from car where car_id = 1
            )
        );

    end;
$$;










--           *********************functions

create or replace function order_payment_share () returns trigger language plpgsql
as
$$
    declare
        car_data record;

    begin

        if old.is_paid = 0 and new.is_paid = 1 then

            select * from car into car_data where car_id = new.car_id;

            insert into order_payment (
                total_price,
                add_tariff_price, 
                investor,
                company, 
                insurance
            ) values (
                new.total_price,
                COALESCE(new.add_tariff_price_day, 0) * new.day_count,
                (car_data.day_price * (CAST(car_data.investor_percentage AS float) / 100)) * new.day_count,
                (car_data.day_price * (1 - (CAST(car_data.investor_percentage AS float) / 100))) * new.day_count,
                car_data.insurance * new.day_count
            );

        end if;

        return new;
    end;
$$;


create trigger trigger_order_payment_share
after insert or update on orders
for each row execute procedure order_payment_share();
