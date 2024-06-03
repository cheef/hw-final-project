CREATE TYPE exception_types AS ENUM ('whitelist', 'blacklist');

CREATE TABLE IF NOT EXISTS exception_lists
(
    id   serial PRIMARY KEY,
    type exception_types NOT NULL,
    cidr varchar(255) NOT NULL
);