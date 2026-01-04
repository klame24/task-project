-- init.sql
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(1000),
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);