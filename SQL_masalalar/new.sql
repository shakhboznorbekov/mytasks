-- create table users(
--     id UUID not null primary key default gen_random_uuid(),
--     first_name varchar(255) not null,
--     last_name varchar(255) not null,
--     nick_name varchar(255) not null,
--     created_at timestamp default current_timestamp,
--     updated_at timestamp default current_timestamp
-- );

-- create table posts(
--     id UUID not null primary key default gen_random_uuid(),
--     user_id UUID not null REFERENCES users(id),
--     created_at timestamp default current_timestamp,
--     updated_at timestamp default current_timestamp);

-- create table chat(
--     id UUID not null primary key default gen_random_uuid(),
--     description varchar(255),
--     deleted boolean,
--     post_id UUID not null REFERENCES posts(id),
--     created_at timestamp default current_timestamp,
--     updated_at timestamp default current_timestamp);

-- INSERT INTO users(first_name,last_name,nick_name)VALUES
-- ('Shaxboz','Norbekov','shekh'),
-- ('Sarvar','Noraliyev','john'),
-- ('Eldor', 'Shomurodov','eldosho'),
-- ('Xasanboy', 'Do''stmatov','xasanchik'),
-- ('Jahongir','Poziljonov', 'bojalar');


-- -- a188b59c-caf0-4cfc-afa5-79258b6fe78c | Shaxboz    | Norbekov   | shekh     | 2022-12-06 19:04:47.706291 | 2022-12-06 19:04:47.706291
-- -- 95ee4d36-8909-4128-b61d-2fbca68e042e | Sarvar     | Noraliyev  | john      | 2022-12-06 19:04:47.706291 | 2022-12-06 19:04:47.706291
-- -- 47f6be11-785f-4fd0-a1c9-57639e87b792 | Eldor      | Shomurodov | eldosho   | 2022-12-06 19:04:47.706291 | 2022-12-06 19:04:47.706291
-- -- 7a4a997e-49cd-4d85-85c1-cd6657f9b94c | Xasanboy   | Do'stmatov | xasanchik | 2022-12-06 19:04:47.706291 | 2022-12-06 19:04:47.706291
-- -- c83524b3-d529-4949-959a-5e3cce837c2f | Jahongir   | Poziljonov | bojalar   | 2022-12-06 19:04:47.706291 | 2022-12-06 19:04:47.706291
-- --(5 rows)


-- INSERT INTO posts(user_id)VALUES
-- ('a188b59c-caf0-4cfc-afa5-79258b6fe78c'),
-- ('95ee4d36-8909-4128-b61d-2fbca68e042e'),
-- ('47f6be11-785f-4fd0-a1c9-57639e87b792'),
-- ('7a4a997e-49cd-4d85-85c1-cd6657f9b94c'),
-- ('c83524b3-d529-4949-959a-5e3cce837c2f');

-- -- 7d05d35d-d5b4-43c3-9d7e-2ae781afe9db | a188b59c-caf0-4cfc-afa5-79258b6fe78c | 2022-12-06 19:09:26.916195 | 2022-12-06 19:09:26.916195
-- -- 14075418-a69c-4af0-91e6-3f712c45403f | 95ee4d36-8909-4128-b61d-2fbca68e042e | 2022-12-06 19:09:26.916195 | 2022-12-06 19:09:26.916195
-- -- 4bfc9dc4-cc1e-4168-a87c-140ce38ed2a8 | 47f6be11-785f-4fd0-a1c9-57639e87b792 | 2022-12-06 19:09:26.916195 | 2022-12-06 19:09:26.916195
-- -- 4fd8ab29-4634-4cef-b8dd-1907782c2ae9 | 7a4a997e-49cd-4d85-85c1-cd6657f9b94c | 2022-12-06 19:09:26.916195 | 2022-12-06 19:09:26.916195
-- -- abd2e693-3a04-4709-be0f-501a6ba5e2ea | c83524b3-d529-4949-959a-5e3cce837c2f | 2022-12-06 19:09:26.916195 | 2022-12-06 19:09:26.916195
-- --(5 rows)



-- INSERT INTO chat(description,deleted,post_id)VALUES
-- ('7d05d35d-d5b4-43c3-9d7e-2ae781afe9db7d05d35d-d5b4-43c3-9d7e-2ae781afe9db', true,'7d05d35d-d5b4-43c3-9d7e-2ae781afe9db'),
-- ('14075418-a69c-4af0-91e6-3f712c45403f14075418-a69c-4af0-91e6-3f712c45403f', true,'14075418-a69c-4af0-91e6-3f712c45403f'),
-- ('4bfc9dc4-cc1e-4168-a87c-140ce38ed2a84bfc9dc4-cc1e-4168-a87c-140ce38ed2a8', true,'4bfc9dc4-cc1e-4168-a87c-140ce38ed2a8'),
-- ('4fd8ab29-4634-4cef-b8dd-1907782c2ae94fd8ab29-4634-4cef-b8dd-1907782c2ae9', true,'4fd8ab29-4634-4cef-b8dd-1907782c2ae9'),
-- ('abd2e693-3a04-4709-be0f-501a6ba5e2eaabd2e693-3a04-4709-be0f-501a6ba5e2ea', true,'abd2e693-3a04-4709-be0f-501a6ba5e2ea'),
-- ('7d05d35d-d5b4-43c3-9d7e-2ae781afe9db7d05d35d-d5b4-43c3-9d7e-2ae781afe9db', true,'7d05d35d-d5b4-43c3-9d7e-2ae781afe9db'),
-- ('7d05d35d-d5b4-43c3-9d7e-2ae781afe9db14075418-a69c-4af0-91e6-3f712c45403f', true,'14075418-a69c-4af0-91e6-3f712c45403f'),
-- ('7d05d35d-d5b4-43c3-9d7e-2ae781afe9db4bfc9dc4-cc1e-4168-a87c-140ce38ed2a8', true,'4bfc9dc4-cc1e-4168-a87c-140ce38ed2a8'),
-- ('7d05d35d-d5b4-43c3-9d7e-2ae781afe9db4fd8ab29-4634-4cef-b8dd-1907782c2ae9', true,'4fd8ab29-4634-4cef-b8dd-1907782c2ae9'),
-- ('7d05d35d-d5b4-43c3-9d7e-2ae781afe9dbabd2e693-3a04-4709-be0f-501a6ba5e2ea', true,'abd2e693-3a04-4709-be0f-501a6ba5e2ea');


-- select
--     user_id,
--     first_name,
--     deleted,
--     c.created_at,
--     description
-- from 
--     users as u
-- left join posts as p on p.user_id=u.id
-- left join chat c on c.post_id=p.id
-- WHERE nick_name ilike 'Shekh';




-- \x gorizontal


