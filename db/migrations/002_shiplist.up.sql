CREATE TABLE shiplist(
    id INTEGER PRIMARY KEY,
    templateId integer NOT NULL,
    level integer  NOT NULL DEFAULT 1,
    exp integer NOT NULL,
    islocked integer NOT NULL,
    skinid integer NOT NULL
);


CREATE TABLE skillidlist(
    skillid integer NOT NULL ,
    skilllvl integer NOT NULL ,
    skillexp integer NOT NULL ,
    shipid integer NOT NULL,

    FOREIGN KEY (shipid)
        REFERENCES shiplist(id)
);
