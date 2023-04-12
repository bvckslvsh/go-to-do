CREATE TABLE users
(
    id serial not null UNIQUE,
    name VARCHAR(255) not null,
    username VARCHAR(255) not null,
    password_hash VARCHAR(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null UNIQUE,
    title VARCHAR(255) not null,
    description VARCHAR(255)
);

CREATE TABLE users_lists
(
    id serial not null UNIQUE,
    user_id int REFERENCES users(id) on delete cascade not null,
    list_id int REFERENCES todo_lists(id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id serial not null UNIQUE,
    title VARCHAR(255) not null,
    description VARCHAR(255),
       done BOOLEAN not null DEFAULT false
);

CREATE TABLE lists_items
(
    id serial not null UNIQUE,
    item_id int REFERENCES todo_items(id) on delete cascade not null,
    list_id int REFERENCES todo_lists(id) on delete cascade not null
);
