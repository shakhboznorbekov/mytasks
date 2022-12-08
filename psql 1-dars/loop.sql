

do 
$$
    DECLARE
        i int =0;
    BEGIN

        LOOP
            raise info '%';

            IF i=10 THEN
                EXIT;
            END IF;

            i = i+1;
        END LOOP;
    END;        

$$;

--for i:=10, i>1; i+=2 {}

do 
$$
    BEGIN
        FOR i IN REVERSE 10..1 BY 2 LOOP
            raise info '%', i;
        END LOOP;
    END;        

$$;



do 
$$
    DECLARE
        movie_data movie%ROWTYPE;
    BEGIN

        FOR movie_data IN (
            SELECT
                *
            FROM movie
        )LOOP
            raise info '%' movie_data;
            END LOOP;
    END;        

$$;




do 
$$
    DECLARE
        movie_data movie%ROWTYPE;
        i int=1;
    BEGIN
        LOOP
            exit WHEN i=10;
            SELECT
                *
            FROM
                movie
            INTO 
                movie_data
            WHERE id =1;
            continue WHEN NOT FOUND;

            raise info '%', movie_data.name;
            i=i+1;
            END LOOP;
    END;        

$$;














--misollar

do
$$
    DECLARE
    SUM int=0;
    BEGIN
        FOR i IN  1..10 LOOP
           SUM=SUM+i;
        END LOOP;
        raise info '%', SUM;
    END;   

$$;


CREATE OR REPLACE FUNCTION Nsum(son int)
RETURNS INT LANGUAGE PLPGSQL
AS
$$
    DECLARE
        SUM INT = 0;
    BEGIN
        FOR i IN  1..son LOOP
           SUM=SUM+i;
        END LOOP;
        RETURN SUM;
    END;   

$$;





CREATE OR REPLACE FUNCTION REVERSENUM(son int)
RETURNS INT LANGUAGE PLPGSQL
AS
$$
    DECLARE
        i int =son;
        reverse int=0;
    BEGIN

        LOOP
            IF i<1 THEN
                EXIT;
            END IF;
            reverse=(reverse*10)+i%10;
            
            i=i/10;
        END LOOP;
        return reverse;
    END;        

$$;



