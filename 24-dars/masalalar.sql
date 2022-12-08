
create table author(
    id serial primary key,
    name_author varchar(24) not null
);


create table genre (
    genre_id serial primary key,
    name_genre  varchar(24)
);


create table book(
    book_id serial primary key,
    title varchar(24),
    price decimal(16,2),
    amount int,
    author_id int not null references author(id),
    genre_id int not null references genre(genre_id)
);


insert into author(name_author) values
('O''tkir Hoshimov'),
('Chingiz Aytmatov'),
('Primqul Qodirov'),
('Said Ahmad'),
('Abdulla Qodiriy');

insert into genre(name_genre) values
('Detektiv',
 'Muhabbat',
 'Sarguzasht',
 'Tarixiy',
 'Hayoti haqida'
);

insert into book (title,price,amount,author_id,genre_id)values
('Dunyoning ishlari',65000,5,1,5),
('Asrga tatigulik kun',70000,10,2,4),
('Yulduzli tunlar',50000,3,3,4),
('Ufq',100000,4,2);
('O''tkan kunlar',75000,5,5);

select 
    * 
from
    book


