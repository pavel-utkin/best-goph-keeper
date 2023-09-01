create table metadata
(
    metadata_id serial PRIMARY KEY,
    entity_id   int        NOT NULL,
    key         varchar(50)  NOT NULL,
    value       varchar(300) NOT NULL,
    type        varchar(50)  NOT NULL,
    created_at  timestamp    NOT NULL,
    updated_at  timestamp    NULL,
    deleted_at  timestamp NULL
);

CREATE INDEX idx_entity_id_key ON metadata (entity_id, key);
CREATE INDEX idx_updated_at ON metadata (updated_at);