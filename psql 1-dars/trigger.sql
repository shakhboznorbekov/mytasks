create table users(
    user_id serial primary key,
    firstname varchar(32) null,
    lastname varchar(32) null,
    middlename varchar(32) null
);


create function trigger_function_fn() returns trigger language plpgsql
AS
$$
    begin

        new.firstname ='Moxirbek';
        return new;

    end;    
$$;


create trigger trigger_users
before insert on users
for each rov execute procedure trigger_function_fn;

insert into users(firstname,lastname,middlename)values('Ravshanbek','Olimov','Akramjonovich');


insert into users(firstname,lastname,middlename)values('Ravshanbek','Olimov','Akramjonovich');






alter table users add column rating int;

create or replace function trigger_users_fn() returns trigger language plpgsql
AS
$$
    begin
        if new.firstname is not null and new.lastname is not null then
            new.rating =5;
        elsif new.middlename is not null and new.lastname is not null then
            new.rating = 3;
        end if;

        return new;
    end;        

$$;



create trigger trigger_users
before insert on users
for each row execute procedure trigger_users_fn();

insert into users(firstname,lastname,middlename) values ('x', 'y','z');
insert into users(lastname,middlename)  values('y','z');