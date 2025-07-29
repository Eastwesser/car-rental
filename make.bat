@echo off
if "%1" == "cars" (
    docker-compose up --build -d
) else if "%1" == "clean" (
    docker-compose down -v
    docker builder prune -af
    docker system prune -af
) else if "%1" == "test" (
    go test ./... -cover
) else (
    echo Usage: make ^<cars^|clean^|test^>
)