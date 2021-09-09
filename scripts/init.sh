#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER smartlyrics WITH ENCRYPTED PASSWORD '$SMARTLYRICS_APP_USER_PASS';
    CREATE DATABASE smartlyrics;
    GRANT ALL PRIVILEGES ON DATABASE smartlyrics TO smartlyrics;
EOSQL
