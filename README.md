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

2. Or try makefiles:
```makefile
make backend
make frontend
make full
```

3. This will run both backend and frontend:

Backend (http://localhost:8080) can be accessed via: [http://localhost:8080/api/cars]

```curl
curl http://localhost:8080/api/cars
curl http://localhost:8080/api/cars?filter=Lamborghini
```

4. Database connection:

```bash
# Check active containers (779793098e50):
docker ps

# Enter postgres:
docker exec -it 779793098e50 psql -U postgres -d carsharing

# List all tables:
\dt
```

SQL
```sql
SELECT * FROM orders ORDER BY created_at DESC LIMIT 5;
```

Frontend can be checked here: [http://localhost:3000]

Enjoy, Go + Angular := fullstack!!! 