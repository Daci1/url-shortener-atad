CREATE SEQUENCE IF NOT EXISTS url_counter
    START 1
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 56800235583  -- 62^6 - 1, max 6-character Base62 short code
    NO CYCLE;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(50) NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS urls (
    id UUID PRIMARY KEY,
    user_id UUID,
    short_url CHAR(6) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- TODO: analytics is redundant since it's calculable from sum of ip_locations visited_counter,
--  decide on this
CREATE TABLE IF NOT EXISTS analytics (
    url_id UUID PRIMARY KEY,
    visited_count INTEGER NOT NULL,
    CONSTRAINT fk_url FOREIGN KEY (url_id) REFERENCES urls (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS unique_visitors (
    url_id UUID PRIMARY KEY,
    visitor_ip INET NOT NULL,
    visited_count INTEGER NOT NULL,
    CONSTRAINT fk_url_visitor FOREIGN KEY (url_id) REFERENCES urls (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS ip_locations (
    city VARCHAR(100),
    region VARCHAR(100),
    country VARCHAR(100),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    url_id UUID NOT NULL,
    visited_counter INTEGER NOT NULL,
    CONSTRAINT fk_url_visitor FOREIGN KEY (url_id) REFERENCES urls (id) ON DELETE CASCADE,
    CONSTRAINT unique_location_per_url UNIQUE (city, region, country, latitude, longitude, url_id)
);
