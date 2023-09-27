
CREATE USER etaireia WITH PASSWORD '12345';

ALTER USER etaireia CREATEDB;

GRANT ALL on tablespace pg_default to etaireia;

CREATE DATABASE signature
    WITH
    OWNER = etaireia
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

