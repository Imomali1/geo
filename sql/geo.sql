DROP TABLE IF EXISTS geo_test.task CASCADE;
DROP TABLE IF EXISTS geo_test.test_point;
DROP SCHEMA IF EXISTS geo_test;

CREATE SCHEMA geo_test;
CREATE TABLE geo_test.test_point (
                                     id integer,
                                     task integer,
                                     x numeric(10, 2),
                                     y numeric(10, 2),
                                     z numeric(10, 2),
                                     PRIMARY KEY (id)
);

CREATE TABLE geo_test.task (
                               id integer PRIMARY KEY,
                               name varchar(256)
);

ALTER TABLE geo_test.test_point
    ADD CONSTRAINT task
        FOREIGN KEY (task)
            REFERENCES geo_test.task (id);

CREATE INDEX index_name ON geo_test.test_point (
                                                id, task
    );

INSERT INTO
    geo_test.task
VALUES
    (1, 'TODO');

INSERT INTO
    geo_test.test_point
VALUES
    (1,1,1,0,0),
    (2,1,0,1,0),
    (3,1,0,0,1);

SELECT * FROM geo_test.test_point LIMIT 1 OFFSET 1;

-- CREATE EXTENSION postgis;

ALTER TABLE geo_test.test_point
    ADD COLUMN geom geometry(PointZ);

UPDATE geo_test.test_point
SET geom = ST_MakePoint(x,y,z);

SELECT ST_AsGeoJSON(geom) FROM geo_test.test_point;