

create table car(
    id serial primary key,
    name varchar(15) not null,
    state_number varchar(11) not null,
    price int not null
);

CREATE TABLE Day(
    id serial primary key,
    car_id int not null REFERENCES car(id)
);


insert into car(name,state_number,price) values
('Spark', '01 A 111 AH', 120000),
('Nexia', '01 B 122 BP', 150000),
('Cobalt', '01 C 123 KK', 170000),
('Lacetti','01 A 007 AA', 200000),
('Malibu', '01 S 571 BP', 300000);


CREATE OR REPLACE FUNCTION getCardata(car_id int) RETURNS TABLE (
    id int,
    name varchar,
    state_number varchar,
    price int
)LANGUAGE PLPGSQL
AS
$$

    BEGIN
        return QUERY
        SELECT
            *
        FROM car
        WHERE car.id = car_id;

    END;
$$;                   




CREATE OR REPLACE FUNCTION getdayCarprice(car_id int, days int) RETURNS TABLE (
    name varchar,
    day int,
    price int
)LANGUAGE PLPGSQL
AS
$$

    BEGIN
        return QUERY
        SELECT
            car.name,
            days,
            car.price*days
        FROM car
        WHERE car.id = car_id;

    END;
$$;    