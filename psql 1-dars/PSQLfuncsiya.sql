SELECT 'Samandar baliq va o''rdak ovini o''rganyapti;

--evaluate

SELECT 2+2;

SELECT $$

Samandar baliq ovini o'rganyapti
2+2


$$;

SELECT $MESSAGE$
    Samandar baliq ovini o'rganyapti$;

    2+2
$message$;


SELECT $$ Uzbekiston $;

--BLOCK SXEMA

--Procedure Language PostgreSQL
--PLPGSQL
DO
$$
    DECLARE
        num INTEGER :=2;
        movieName varchar;
        movieTable RECORD;

    BEGIN
        --SELECT
        --  name
        --FROM movie
        -- INTO movieName
        --WHERE id = 2
        --;
        --raise info '%', movieTable;
        raise info '%', movieTable;
    END;
$$;

--FUNCTION

CREATE FUNCTION getX() RETURNS INT LANGUAGE PLPGSQL
AS
$$
    DECLARE
        num INTEGER =10;
    BEGIN

        num= num +1;
        return num;
    END;
$$;


SELECT getX();
/df --function show

CREATE OR REPLACE FUNCTION getMovieName(movie_id INT)
RETURNS VARCHAR LANGUAGE PLPGSQL
AS
$$
    DECLARE
        movieName varchar;
    BEGIN

        SELECT
            m.name
        FROM movie AS m 
        INTO movieName
        WHERE m.id=movie_id;
        return movieName;
    END;
$$;               





CREATE OR REPLACE FUNCTION getMovieName(movie_id INT)
RETURNS VARCHAR LANGUAGE PLPGSQL
AS
$$
    DECLARE
        movieName varchar;
    BEGIN

        SELECT
            m.name
        FROM movie AS m 
        INTO movieName
        WHERE m.id=movie_id;
        return movieName;
    END;
$$;    

get

create table movie(
    id serial primary key,
    name varchar(25) not null,
    data_create int
);




insert into movie(name, data_create)values
('Shaytanat',2002),
('Avatar 1', 2010),
('Qog''oz bino',2022);



create table comments(
    id serial primary key,
    comment varchar(50),
    movie_id int REFERENCES movie(id)
);


insert into comments(comment, movie_id)values
('yaxshi kino ekan',1),
('oxirida yomon holatda tugagani yoqmadi',2),
('qachon ikkinchi qismi chiqarkana',2),
('ajoyib zor ekan',3),
('ajoyib fantastika',2),
('eng yaxshi o''zbek kinosi shu bo''lsa kerak',1);






create table 


CREATE OR REPLACE FUNCTION get_users(son int,soz varchar)
RETURNS INT LANGUAGE PLPGSQL
AS
$$
   
    BEGIN

        SELECT 
            *
        FROM
            letters AS l1
        WHERE
        (
            SELECT
                count(*)%2<>0
            FROM
                letters AS l2
            WHERE l2.id<l1.id        
        )        
    END;        

$$;

create table userss(
    id serial primary key,
    name varchar(15)
);




SELECT 
    *
FROM
    letters AS l1
WHERE
(
    SELECT
        count(*)%2<>0
    FROM
        letters AS l2
    WHERE l2.id<l1.id        
)        




create table element(
    
)