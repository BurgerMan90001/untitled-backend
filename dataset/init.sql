
CREATE DATABASE untitled;

CREATE TABLE users (
    _id TEXT,
    username TEXT,
    email TEXT,
)

COPY users
FROM '/docker-entrypoint-initdb.d/users.csv'
DELIMITER ','
CSV HEADER;