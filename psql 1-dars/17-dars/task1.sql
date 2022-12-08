create table cource1(
    id serial primary key,
    name varchar(20),
    students varchar[]
);


insert into cource1(name, students) values
('Go', '{"Samandar","Ali"}');


select 
    name,
    array_length(students,1) as a,
    students
from
    cource1;