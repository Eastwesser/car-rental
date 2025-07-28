# Car Rental App

## Project Structure

- `backend/`: Go server with API endpoints
- `frontend/`: Angular frontend application
- `docker-compose.yaml`: Docker configuration

## Setup

1. Build and run with Docker:
```bash
# full project
docker-compose up --build
# only backend
docker-compose up --build backend
# only frontend
docker-compose up --build frontend

# CLEANUP
docker-compose down -v
docker builder prune -af
docker system prune -af
```

2. Or try makefiles!
```makefile
make backend
make frontend
make full
```

3. This will run both backend and frontend.

Backend (http://localhost:8080) can be accessed via: [http://localhost:8080/api/cars]

```curl
curl http://localhost:8080/api/cars
curl http://localhost:8080/api/cars?filter=BMW
```

Frontend (http://localhost:3000) can be checked here: [http://localhost]

Enjoy, Go + Angular := fullstack!!! 