
CREATE DATABASE untitled;

CREATE TABLE users (
    id INT PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    email TEXT,
)

COPY users
FROM '/docker-entrypoint-initdb.d/users.csv'
DELIMITER ','
CSV HEADER;