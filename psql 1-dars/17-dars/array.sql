create table car (
    id serial primary key,
    name varchar(24),
    colors varchar[]
);

insert into car(name, colors) values
('Mercedes', '{"Black", "White"}');


insert into car(name, colors) values
('Toyoto', '{"Green", "Black"}');


insert into car(name, colors) values
('Audi', '{"White", "Choco"}');

select
    name,
    colors[1]
from 
    car;


select
    name,
    colors[1]
from 
    car
where 'Black'=any(colors);


update
    car
set
    colors=['brown']
where id =3;

update 
    car
set
    colors=array_append(colors,'Choco')
where id = 3;


select array_append(array['1','2'], '3');

select array_append('3', array['1','2']);

select array_remove(array['A','B','C'], 'C');