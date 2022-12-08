create table movie(
    id serial primary key,
    name varchar(255),
    stars int
);


create table clients(
    id serial primary key,
    first_name varchar(255),
    last_name varchar(255),
    movie_id int not null references movie(id)
);


insert into movie (name,stars) values
('Avatar',4),
('G''aroyib juftlik',2),
('Qasoskorlar', 5)
;

insert into clients(first_name,last_name,movie_id)values
('Shaxboz','Norbekov',3),
('Sarvar','Noraliyev',2),
('Bobur', 'Ochulov',1);

select 
    clients.first_name,
    clients.last_name,
    movie.name

from 
    movie    
join clients on movie.id=clients.id
where clients.id=2;