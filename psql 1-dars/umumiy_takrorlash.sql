

create table car (
    car_id serial primary key,
    state_number varchar,
    status varchar default 'in_stock',
    day_price int not null,
    km_price int not null,
    km_limit int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);


create table client (
    client_id serial primary key,
    name varchar
);


create table orders (
    order_id serial primary key,
    price decimal,
    status varchar not null,
    day_count int,
    give_km int not null,
    recieve_km int,
    car_id int not null references car(car_id),
    client_id int not null references client(client_id),
    created_at timestamp default current_timestamp
);

create table order_cost(
    order_cost_id serial primary key,
    price decimal,
    order_id int references orders(order_id)
    
);


insert into car (state_number, status, day_price, km_price, km_limit, updated_at) values
('01 A 111 AA', 'in_stock',500000,1000,200, now()),
('01 W 556 PB', 'in_stock', 400000,2000,300,now()),
('11 111 YYY', 'in_stock', 300000,1500,150,now()),
('77 Z 777 ZZ', 'in_stock', 250000,1000,200,now()),
('01 G 900 MA', 'in_stock', 200000,2000, 300,now());

insert into client (name) values
('Samandar'),
('Moxirbek'),
('Jahongir');

insert into orders (status, day_count, give_km, car_id, client_id) values
('new', 5, 2000, 4, 1),
('new', 7, 1500, 5, 2),
('new', 6, 1100, 1, 3);

update orders set
    price = 3000000,
    recieve_km=5000,
    status = 'client_returned'
where order_id = 2;

--2-vazifa
create trigger calculate
after insert or update on orders
for each row execute procedure trigger_order_status1();


create or replace function trigger_order_status1() returns trigger language plpgsql
as
$$
    begin
       
        if new.recieve_km - new.give_km > new.day_count * 200 then
            insert into order_cost(order_id,price) values (new.order_id,(new.recieve_km - (new.give_km + new.day_count * 200 ) * 1000));
        end if;

        return new;
    end;
$$;


-- create or replace function trigger_order_status1() returns trigger language plpgsql
-- as
-- $$
--     begin
       
--         if new.recieve_km - new.give_km > new.day_count * car.km_limit then
--             insert into order_cost(order_id,price) values (new.order_id,(new.recieve_km - (new.give_km + new.day_count * car.km_limit ) * car.km_price));
--         end if;

--         return new;
--     end;
-- $$;










1. orders kun qoshasiz
    day_count = 5
    give_km = 2000
    receive_km = 3500

   car kunlik tolov qoshasiz
    day_price = 500000
    km_price = 1000
    km_limit = 200

2. order_cost table qilasiz
    order_id
    price


create table order_cost (
    order_cost_id serial not null primary key,
    price int not null,
    order_id int not null references orders(order_id)
);


3. calculate trigger yaratasiz
    zakaz qabul qilinganda kunlik km hisoblanib oshiqchasiga qarz sifatida
    order_cost table ga insert qiling.


create or replace function pere_limit() returns trigger language plpgsql
as
$$
    declare
        actual_km int = 0;
        diff int = 0;
        car_km_price int = 0;
    begin

        if old.status = 'client_took' and new.status = 'client_returned'
            then

                select
                    new.day_count * c1.km_limit,
                    c1.km_price
                from
                (
                    select
                        c.km_limit,
                        c.day_price,
                        c.km_price
                    from
                        car as c
                    where c.car_id = new.car_id
                ) as c1
                into actual_km, car_km_price;

                diff = (new.receive_km - new.give_km) - actual_km;

                if diff > 0 then
                    insert into order_cost (price, order_id) values
                    (diff * car_km_price, new.order_id);
                end if;

        end if;

        return new;
    end;
$$;

create trigger tigger_pere_limit
after insert or update on orders
for each row execute procedure pere_limit();


insert into orders (price, status, car_id, client_id, day_count) values
((select day_price * 5 from car where car_id = 2), 'new', 2, 2, 5);

update orders set status = 'client_took', give_km = 20000 where order_id = 4;

update orders set status = 'client_returned', receive_km = 21500 where order_id = 4;












--4-vazifa
select 
    cl.name, 
    oc.price 
from order_cost  as oc
left JOIN orders as o On oc.order_id = o.order_id
left JOIN client as cl On o.client_id = cl.client_id;



4. qarzga kirgan clientning royxatini chiqazing
    client   | price
    samandar | 500000

select
    c.name,
    oc.price
from
    order_cost as oc
join orders as o on o.order_id = oc.order_id
join client as c on c.client_id = o.client_id;
































-- create or replace function trigger_order_status() returns trigger language plpgsql
-- as
-- $$
--     begin




--         if new.status = 'new' then
--             update car set 
--                 status = 'in_booked',
--                 updated_at = now()
--             where car_id = new.car_id;
--         elsif old.status = 'new' and new.status = 'client_took' then
--             update car set
--                 status = 'in_use',
--                 updated_at = now()
--             where car_id = new.car_id;
--         elsif old.status = 'client_took' and new.status = 'client_returned' then
--             update car set 
--                 status = 'in_stock',
--                 updated_at = now()
--             where car_id = new.car_id;
--         end if;

--         return new;
--     end;
-- $$;



-- create trigger trigger_orders
-- after insert or update on orders
-- for each row execute procedure trigger_order_status();












-- // cars tabledan new.car_id orqali km_limit, km_price

-- // new.recieve_km - new.give_km > new.day_count * km_limit
-- // insert into order_cost (new.id, (new.recieve_km - new.give_km) * km_price)
       


 -- if new.recieve_km is not null then
        --     select 
        --         km_limit,
        --         km_price
        --     from car;
        -- end if;