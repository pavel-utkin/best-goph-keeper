create table session
(
    session_id       serial PRIMARY KEY,
    user_id          int         NOT NULL references users (user_id) on delete cascade,
    session_key      varchar(30) NOT NULL,
    created_at       timestamp   NOT NULL,
    end_date_at       timestamp   NOT NULL
);