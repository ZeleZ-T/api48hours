USE database;

CREATE TABLE IF NOT EXISTS users (
    email VARCHAR(25) PRIMARY KEY,
    password VARCHAR(500) NOT NULL
);