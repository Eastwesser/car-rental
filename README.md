# Car Rental App

## Project Structure

- `backend/`: Go server with API endpoints
- `frontend/`: Angular frontend application
- `docker-compose.yaml`: Docker configuration

## Ports

- `Frontend`: carsharing-go-frontend-1 (port 3000)
- `Backend`: carsharing-go-backend-1 (port 8080)
- `Database`: carsharing-go-db-1 (port 5432)

## Routes

- `Get cars`: api/v1/getcars
- `Make order`: api/v1/order

## Setup

1. Build and run with Docker:
```bash
# full project
docker-compose up --build

# Or try makefiles:
make cars

# CLEANUP
docker-compose down -v
docker builder prune -af
docker system prune -af
```

2. `make cars` will run both backend and frontend:

- Backend can be accessed via: [http://localhost:8080/api/cars]

- Frontend can be checked here: [http://localhost:3000]

3. cURL:

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

5. SQL:

```sql
SELECT * FROM orders ORDER BY created_at DESC LIMIT 5;
```

## Tests:

```bash
# Run tests with coverage and save to file
go test ./... -coverprofile=coverage.out

# View coverage report in terminal
go tool cover -func=coverage.out

# For a quick summary:
go test ./... -cover
```

## DockerHub updates:

Run these commands from root directory where docker-compose is:

```bash
# Login to Docker Hub
docker login

# Build and push backend
docker build -t eastwesser/carsharing-backend -f backend/Dockerfile .
docker push eastwesser/carsharing-backend

# Build and push frontend
docker build -t eastwesser/carsharing-frontend -f frontend/Dockerfile ./frontend
docker push eastwesser/carsharing-frontend
```

Enjoy, Go (Chi) + TS (Angular) := fullstack!!! 