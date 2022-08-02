#!/bin/bash
set -e
echo "Im here!"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER logisticapp WITH PASSWORD 'vn53nag';
	CREATE DATABASE logisticapp;
	GRANT ALL PRIVILEGES ON DATABASE logisticapp TO logisticapp;
	ALTER USER logisticapp REPLICATION SUPERUSER;

EOSQL