create table users
(
    user_id    serial PRIMARY KEY,
    username   varchar(25) NOT NULL,
    password   varchar(30) NOT NULL,
    created_at timestamp   NOT NULL,
    deleted_at timestamp NULL
);