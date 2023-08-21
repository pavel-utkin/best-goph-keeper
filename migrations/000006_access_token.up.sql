create table access_token
(
    access_token     varchar(32) NOT NULL,
    user_id          int         NOT NULL references users (user_id) on delete cascade,
    created_at       timestamp   NOT NULL,
    end_date_at      timestamp   NOT NULL
);

CREATE UNIQUE INDEX idx_unique_access_token ON access_token (access_token);