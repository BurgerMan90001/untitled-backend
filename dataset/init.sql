
CREATE DATABASE IF NOT EXISTS untitled;

CREATE TABLE IF NOT EXISTS users (
    id TEXT
    username TEXT
    email TEXT
)

COPY users
FROM '/docker-entrypoint-initdb.d/users.csv'
DELIMITER ','
CSV HEADER;