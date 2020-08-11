-- shortener DB
CREATE TABLE IF NOT EXISTS links
(
    id          VARCHAR(150) PRIMARY KEY,
    link        TEXT NOT NULL UNIQUE,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);