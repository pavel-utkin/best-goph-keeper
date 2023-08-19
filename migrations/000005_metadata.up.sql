create table metadata
(
    metadata_id      serial PRIMARY KEY,
    cart_name        varchar(50) NOT NULL,
    cart_description varchar(300) NULL,
    created_at       timestamp   NOT NULL,
    updated_at       timestamp   NOT NULL,
    deleted_at       timestamp NULL
);