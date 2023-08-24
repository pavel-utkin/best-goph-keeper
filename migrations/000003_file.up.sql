create table file
(
    file_id        serial PRIMARY KEY,
    user_id        int          NOT NULL references users (user_id) on delete cascade,
    file_name      varchar(100) NOT NULL,
    file_path      varchar(200) NOT NULL,
    file_extension varchar(10)  NOT NULL,
    created_at     timestamp    NOT NULL,
    updated_at     timestamp    NOT NULL,
    deleted_at     timestamp NULL
);