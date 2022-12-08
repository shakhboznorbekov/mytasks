create table tables(
    table_id serial primary key,
    number varchar(10) not null
);

insert into tables(number)values
('1-stol'),
('2-stol'),
('3-stol'),
('4-stol'),
('5-stol'),
('6-stol'),
('7-stol'),
('8-stol'),
('9-stol'),
('10-stol');










create table orders(
    order_id serial primary key,
    table_id int references tables(table_id),
    created_at timestamp default current_timestamp,
    closed_at timestamp
);

insert into orders(table_id,created_at,closed_at)values
    (1,'2022-11-10','2022-11-10'),
    (2,'2022-9-10','2022-9-10'),
    (3,'2022-10-10','2022-10-10'),
    (4,'2022-8-10','2022-8-10'),
    (5,'2022-9-10','2022-9-10'),
    (6,'2022-10-15','2022-10-15');









create table products(
    product_id serial primary key,
    name varchar(64),
    prise decimal(16, 2)
);

insert into products (name, price) values
('osh', 28000.00),
('mastava', 25000.00),
('manti', 5000.00),
('choy', 2000.00),
('Limon choy', 6000.00),
('sezar salat', 7000.00),
('pentuza', 7000.00);








create table components(
    component_id serial primary key,
    name varchar(24)
);

insert into components(name)values
('sabzi'),
('yog'),
('guruch'),
('gosht'),
('kartoshka'),
('piyoz'),
('suv'),
('un');









create table ingredients(
    ingredient_id serial primary key,
    product_id int not null references products(product_id),
    component_id int not null references components(component_id),
    weight int2

);

insert into ingredients(product_id,component_id,weight)values
(1,1,4),
(1,2,1),
(1,3,1),
(1,4,1),
(1,5,2),
(1,6,1),
(1,7,6);








create table order_details (
    order_detail_id serial primary key,
    quantity int2 not null,
    order_id int not null references orders(order_id),
    product_id int not null references products(product_id)
);


insert into order_details(quantity,order_id,product_id) values
    (1,1,1),
    (2,2,1),
    (3,3,4),
    (4,2,2),
    (5,5,1),
    (6,7,1); 









create table types(
    type_id serial primary key,
    name varchar(64)
);


insert into types (name) values
('suyuq'),
('quyuq'),
('ichimliklar'),
('salat');









create table categories(
    category_id serial primary key,
    name varchar(64)
);

insert into categories (name) values
('Milliy taomlar'),
('Yevropa taomlar'),
('Xitoy taomlar');










create table type_register (
    type_register_id serial primary key,
    category_id int not null references categories(category_id),
    type_id int not null references types(type_id),
    product_id int not null references products(product_id)
);

insert into type_register(category_id, type_id, product_id) values
(1, 2, 1),
(1, 1, 2),
(1, 3, 4),
(2, 4, 7),
(3, 4, 6),
(1, 2, 3),
(1, 3, 5);





select
    c.name,
    ARRAY_AGG(DISTINCT t.name),
    ARRAY_AGG(p.name)
from
    categories AS c
JOIN type_register AS tr USING(category_id)
JOIN types AS t USING(type_id)
JOIN products AS p USING(product_id)
GROUP BY c.name
;
--  2 - task 
select
    count(closed_at)
from 
    orders
JOIN order_details USING(order_id)
WHERE order_details.product_id = 1 AND orders.closed_at = '2022-11-10';



