

create table product_groups (
    group_id serial primary key,
    group_name varchar(255) not null
);

create table products (
    product_id serial primary key,
    product_name varchar(255) not null,
    price decimal(11, 2),
    group_id int not null references product_groups(group_id)
);

insert into product_groups (group_name) 
values
('Smartphone'),
('Laptop'),
('Tablet');


insert into products (product_name, price, group_id)
values
('Iphone', 1400000, 1),
('Acer', 1200000, 2),
('Samsung A2', 800000, 3),
('Nokia 1280', 200000, 1),
('HP', 900000, 2),
('Ipad', 1300000, 3),
('Oppo', 1000000, 1),
('Xiomi', 700000, 3),
('Sony', 680000, 3),
('Lenovo Thinkpad', 980000, 2);


select
    group_name,
    product_name,
    price,
    AVG(price) OVER(
        partition by group_name
    )
from
    products
natural join product_groups
;

select
    count(*) over(
        partition by group_name
    ),
    group_name,
    product_name,
    price
from
    products
natural join product_groups
;

select
    group_name,
    product_name,
    price,
    ROW_NUMBER() OVER(
        partition by group_name
        order by price
    )
from
    products
natural join product_groups
;

select
    group_name,
    product_name,
    price,
    first_value(price) OVER(
        partition by group_name
        order by price desc
    )
from
    products
natural join product_groups
;