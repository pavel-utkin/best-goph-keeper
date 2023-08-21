create table text
(
    text_id    serial PRIMARY KEY,
    user_id    int       NOT NULL references users (user_id) on delete cascade,
    text       bytea     NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
);