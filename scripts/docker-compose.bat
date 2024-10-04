@echo off
cd %cd%\..\

set DEBUG=false
docker compose --env-file ./.env up