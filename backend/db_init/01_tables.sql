CREATE SEQUENCE IF NOT EXISTS url_counter
    START 1
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 56800235583  -- 62^6 - 1, max 6-character Base62 short code
    NO CYCLE;

CREATE TABLE IF NOT EXISTS urls (
    id UUID PRIMARY KEY,
    short_url CHAR(6) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
