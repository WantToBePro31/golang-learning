CREATE TABLE persons (
    id INTEGER primary key,
    NIK VARCHAR(255) not null UNIQUE,
    fullname VARCHAR(255) not null,
    gender VARCHAR(50) not null,
    birth_date DATE not null,
    is_married BOOLEAN,
    height FLOAT,
    weight FLOAT,
    address TEXT
);