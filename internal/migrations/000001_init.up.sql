CREATE TABLE IF NOT EXISTS users
(
    id       serial PRIMARY KEY,
    username varchar(255) not null unique,
    password varchar(255) not null,
    role     varchar(255) not null default 'user'
);

CREATE TABLE IF NOT EXISTS products
(
    id          serial PRIMARY KEY,
    name        text[] not null,
    description text[],
    price       int    not null,
    discount    int
);

CREATE TABLE IF NOT EXISTS contacts
(
    id          serial PRIMARY KEY,
    coordinates text       not null,
    phone       varchar(255) not null,
    vip_phone   varchar(255) not null,
    address     text       not null
);

CREATE TABLE IF NOT EXISTS categories
(
    id   serial PRIMARY KEY,
    name text[] not null
);

CREATE TABLE IF NOT EXISTS category_items
(
    id          serial PRIMARY KEY,
    category_id int references categories (id) on delete cascade not null,
    product_id  int references products (id) on delete cascade   not null
);
