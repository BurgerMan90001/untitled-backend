
CREATE DATABASE untitled;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT,
    email TEXT,
)

COPY users
FROM '/docker-entrypoint-initdb.d/users.csv'
DELIMITER ','
CSV HEADER;