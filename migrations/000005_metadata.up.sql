create table metadata
(
    metadata_id serial PRIMARY KEY,
    name        varchar(50) NOT NULL,
    description varchar(300) NULL,
    created_at  timestamp   NOT NULL,
    updated_at  timestamp   NOT NULL,
    deleted_at  timestamp NULL
);