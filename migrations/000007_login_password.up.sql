create table login_password
(
    login_password_id     serial PRIMARY KEY,
    user_id     int       NOT NULL references users (user_id) on delete cascade,
    data        bytea     NOT NULL,
    created_at  timestamp NOT NULL,
    updated_at  timestamp NOT NULL,
    deleted_at  timestamp NULL
);