@echo off
set BASE_PATH=%~dp0backend

for %%F in (auth_service sekolah sc-service) do (
    start cmd /k "cd /d %BASE_PATH%\%%F && go run main.go"
)
