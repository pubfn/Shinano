CREATE TABLE commander (
    id integer NOT NULL,
    name text NOT NULL,
    level integer NOT NULL DEFAULT 1,
    exp integer NOT NULL
);


CREATE TABLE resourcelist (
    type integer NOT NULL,
    num integer NOT NULL
)
