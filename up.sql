-- shortener DB
CREATE TABLE IF NOT EXISTS bloods
(
    id          VARCHAR(150) PRIMARY KEY,
    link        TEXT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);