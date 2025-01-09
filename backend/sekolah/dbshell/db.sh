#!/bin/bash
set -e
export PGPASSWORD=postgres123;
psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "pendataan" <<-EOSQL
    CREATE DATABASE pendataan;
    GRANT ALL PRIVILAGES ON DATABASE pendataan TO "postgres";
EOSQL