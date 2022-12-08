
CREATE TABLE product (
    id UUID NOT NULL PRIMARY KEY default gen_random_uuid(),
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE customer (
    id UUID NOT NULL PRIMARY KEY default gen_random_uuid(),
    name VARCHAR NOT NULL,
    balance NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE customer_cost (
    id UUID NOT NULL PRIMARY KEY default gen_random_uuid(),
    customer_id UUID NOT NULL REFERENCES customer(id),
    product_id UUID NOT NULL REFERENCES product(id),
    price NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cashier (
    id UUID NOT NULL PRIMARY KEY default gen_random_uuid(),
    name VARCHAR NOT NULL
);

CREATE TABLE withdraw (
    id UUID NOT NULL PRIMARY KEY default gen_random_uuid(),
    total_summa NUMERIC NOT NULL,
    cashier_id UUID NOT NULL REFERENCES cashier(id),
    customer_id UUID NOT NULL REFERENCES customer(id)
);
INSERT INTO product(name, price) VALUES
('smartphone', round(random() * 100000)),
('laptop', round(random() * 100000)),
('fragrance', round(random() * 100000)),
('skincare', round(random() * 100000)),
('grocerie', round(random() * 100000)),
('home-decoration', round(random() * 100000)),
('furniture', round(random() * 100000)),
('top', round(random() * 100000)),
('womens-dresse', round(random() * 100000)),
('womens-shoe', round(random() * 100000)),
('mens-shirt', round(random() * 100000)),
('mens-shoe', round(random() * 100000)),
('mens-watche', round(random() * 100000)),
('womens-watche', round(random() * 100000)),
('womens-bag', round(random() * 100000)),
('womens-jewellery', round(random() * 100000)),
('sunglasse', round(random() * 100000)),
('automotive', round(random() * 100000)),
('motorcycle', round(random() * 100000)),
('lighting', round(random() * 100000));


INSERT INTO customer(name, balance) VALUES
('Samandar', 600000),
('Moxirbek', 20600000),
('Shaxboz', 12402100),
('Jahongir', 6400650),
('Ravshanbek', 9000000);

INSERT INTO cashier(name) VALUES
('Anna'),
('Tom');

INSERT INTO customer_cost(customer_id, product_id, price) VALUES
('3d8cd431-1063-410a-a80a-9a912b318dbf', 'fe2397a9-fe43-4dea-a715-d76411d7483e', 46313),
('8fd6f7c3-a7ce-4cad-9570-a1a851492c0c', 'fe2397a9-fe43-4dea-a715-d76411d7483e', 46313),
('8fd6f7c3-a7ce-4cad-9570-a1a851492c0c', 'd7e06d1f-5677-476f-97c3-85bee5704388', 99986),
('9c057e38-94fa-4b71-af82-70a683959e68', 'd7e06d1f-5677-476f-97c3-85bee5704388', 99986),
('20206208-2112-4048-a5df-052db59fb816', '88174dd6-894a-4691-9474-746e0eb0981a', 45645),
('20206208-2112-4048-a5df-052db59fb816', '4b6eb38a-1b51-47d8-a7c9-97c42229b918', 89541),
('20206208-2112-4048-a5df-052db59fb816', 'c738aa8a-50f5-4e5c-965c-53738ddff8c3', 45612),
('20206208-2112-4048-a5df-052db59fb816', 'f08e8630-0b6f-49ba-b1cc-0f50ceaee770', 98765),
('9c057e38-94fa-4b71-af82-70a683959e68', '88174dd6-894a-4691-9474-746e0eb0981a', 45645),
('9c057e38-94fa-4b71-af82-70a683959e68', '4b6eb38a-1b51-47d8-a7c9-97c42229b918', 89541),
('9c057e38-94fa-4b71-af82-70a683959e68', 'c738aa8a-50f5-4e5c-965c-53738ddff8c3', 45612),
('9c057e38-94fa-4b71-af82-70a683959e68', 'f08e8630-0b6f-49ba-b1cc-0f50ceaee770', 98765),
('9c057e38-94fa-4b71-af82-70a683959e68', '27195370-dc1f-4faa-9e8a-c156874decb1', 45689),
('9c057e38-94fa-4b71-af82-70a683959e68', '77a63a7e-d9e5-472c-aff8-4ff798eb0f1b', 78932),
('9c057e38-94fa-4b71-af82-70a683959e68', '5888560d-cc33-425a-b3f4-780096129558', 32165),
('20206208-2112-4048-a5df-052db59fb816', '27195370-dc1f-4faa-9e8a-c156874decb1', 45689),
('3d8cd431-1063-410a-a80a-9a912b318dbf', '27195370-dc1f-4faa-9e8a-c156874decb1', 45689),
('20206208-2112-4048-a5df-052db59fb816', '77a63a7e-d9e5-472c-aff8-4ff798eb0f1b', 78932),
('20206208-2112-4048-a5df-052db59fb816', '5888560d-cc33-425a-b3f4-780096129558', 32165),
('20206208-2112-4048-a5df-052db59fb816', 'fe2397a9-fe43-4dea-a715-d76411d7483e', 46313),
('e801e802-31dd-4aad-a8e3-0a6b57f5c92e', '27195370-dc1f-4faa-9e8a-c156874decb1', 45689),
('e801e802-31dd-4aad-a8e3-0a6b57f5c92e', '77a63a7e-d9e5-472c-aff8-4ff798eb0f1b', 78932),
('e801e802-31dd-4aad-a8e3-0a6b57f5c92e', '5888560d-cc33-425a-b3f4-780096129558', 32165);



INSERT INTO withdraw(total_summa, cashier_id, customer_id) VALUES
(
    (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '3d8cd431-1063-410a-a80a-9a912b318dbf'
    ),
    '6b6df294-7a0c-44ca-8661-b8cb630062c1',
    '3d8cd431-1063-410a-a80a-9a912b318dbf'
);

update customer set balance = balance - (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '3d8cd431-1063-410a-a80a-9a912b318dbf'
    ) WHERE id = '3d8cd431-1063-410a-a80a-9a912b318dbf'






--2chisi
INSERT INTO withdraw(total_summa, cashier_id, customer_id) VALUES
(
    (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '465d7b6b-d577-4d2c-9044-95db08bcbc0f'
    ),
    '6b6df294-7a0c-44ca-8661-b8cb630062c1',
    '465d7b6b-d577-4d2c-9044-95db08bcbc0f'
);

update customer set balance = balance - (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '465d7b6b-d577-4d2c-9044-95db08bcbc0f'
    ) WHERE id = '465d7b6b-d577-4d2c-9044-95db08bcbc0f'



--3chisi
INSERT INTO withdraw(total_summa, cashier_id, customer_id) VALUES
(
    (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '9c057e38-94fa-4b71-af82-70a683959e68'
    ),
    '6b6df294-7a0c-44ca-8661-b8cb630062c1',
    '9c057e38-94fa-4b71-af82-70a683959e68'
);

update customer set balance = balance - (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '9c057e38-94fa-4b71-af82-70a683959e68'
    ) WHERE id = '9c057e38-94fa-4b71-af82-70a683959e68'




--4chisi
INSERT INTO withdraw(total_summa, cashier_id, customer_id) VALUES
(
    (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '8fd6f7c3-a7ce-4cad-9570-a1a851492c0c'
    ),
    '6b6df294-7a0c-44ca-8661-b8cb630062c1',
    '8fd6f7c3-a7ce-4cad-9570-a1a851492c0c'
);

update customer set balance = balance - (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = '8fd6f7c3-a7ce-4cad-9570-a1a851492c0c'
    ) WHERE id = '8fd6f7c3-a7ce-4cad-9570-a1a851492c0c'



--5chisi
INSERT INTO withdraw(total_summa, cashier_id, customer_id) VALUES
(
    (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = 'e801e802-31dd-4aad-a8e3-0a6b57f5c92e'
    ),
    '6b6df294-7a0c-44ca-8661-b8cb630062c1',
    'e801e802-31dd-4aad-a8e3-0a6b57f5c92e'
);

update customer set balance = balance - (
        SELECT
            SUM(price)
        FROM customer_cost
        WHERE customer_id = 'e801e802-31dd-4aad-a8e3-0a6b57f5c92e'
    ) WHERE id = 'e801e802-31dd-4aad-a8e3-0a6b57f5c92e'







SELECT
    c.name,
    csh.name,
    w.total_summa,
    TO_CHAR(w.created_at, 'YYYY-MM-DD HH24-MI-SS')
FROM
    customer AS c
INNER JOIN withdraw AS w ON w.customer_id = c.id
INNER JOIN cashier AS csh ON csh.id = w.cashier_id



ALTER TABLE withdraw ADD COLUMN created_at timestamp default current_timestamp;


--2-shart
select 
    Sum(total_summa)
FROM
    withdraw;


--3-shart
SELECT
    count(p.id),
    p.name
FROM
    customer_cost AS cc
INNER JOIN product AS p ON p.id = cc.product_id
group by p.name
order by count(p.id) desc limit 1;


    