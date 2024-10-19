CREATE TABLE demo_users_table (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name    TEXT NOT NULL,
    last_name     TEXT NOT NULL,
    resume        TEXT,
    birth_date    DATE,
    password      TEXT NOT NULL,
    terms         BOOLEAN NOT NULL
);