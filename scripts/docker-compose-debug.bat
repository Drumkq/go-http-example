@echo off
cd %cd%\..\

set DEBUG=true
docker compose --env-file ./.env.development up