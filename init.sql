CREATE DATABASE IF NOT EXISTS map_database;
USE map_database;

CREATE TABLE IF NOT EXISTS users (
    email VARCHAR(25) PRIMARY KEY,
    password VARCHAR(500) NOT NULL
);

CREATE TABLE IF NOT EXISTS map_params
(
    id           INTEGER PRIMARY KEY AUTO_INCREMENT,
    seed         INTEGER NOT NULL,
    width        INTEGER NOT NULL,
    height       INTEGER NOT NULL,
    smoothness   float   NOT NULL,
    water_level  float   NOT NULL,
    perlin_noise boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS maps
(
    name  VARCHAR(50) PRIMARY KEY,
    owner VARCHAR(25) NOT NULL,
    FOREIGN KEY (owner) REFERENCES users (email),
    params_id INTEGER NOT NULL,
    FOREIGN KEY (params_id) REFERENCES map_params (id)
);