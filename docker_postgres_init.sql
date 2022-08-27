CREATE USER mtspremium WITH PASSWORD 'my-secret-pw' CREATEDB;
CREATE DATABASE mtspremium
    WITH 
    OWNER = mtspremium
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;