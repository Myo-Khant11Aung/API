-- migrate:up
CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(225) NOT NULL
);

-- migrate:down
DROP TABLE tasks;
