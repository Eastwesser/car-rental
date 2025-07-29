# Car Rental App

Premium car sharing service offering exclusive vehicles from Lamborghini to Ferrari. 
Book your dream car in 3 simple steps.

## System Requirements
- Docker 20.10+
- Docker Compose 2.0+
- (Optional for local development):
  - Go 1.24+ (for backend development)
  - Node.js 20+ and Angular CLI (for frontend development)

## Windows Setup Guide

1. **Install Prerequisites**:
   - Install [Docker Desktop](https://www.docker.com/products/docker-desktop/)
   - Enable WSL2 backend (recommended) during installation

2. **First-Time Setup**:

```bash
# Clone the repository
git clone https://github.com/eastwesser/carsharing-go.git
cd carsharing-go
``` 

## Project Structure

- `backend/`: Go server with API endpoints
- `frontend/`: Angular frontend application
- `docker-compose.yaml`: Docker configuration

## Ports

- `Frontend`: carsharing-go-frontend-1 (port 3000)
- `Backend`: carsharing-go-backend-1 (port 8080)
- `Database`: carsharing-go-db-1 (port 5432)

## API Endpoints (Routes For Frontend Developers)

- Base URL: http://localhost:8080/api

| Method    | Endpoint       | Parameters           | Description                          | Example                          |
|-----------|----------------|----------------------|--------------------------------------|----------------------------------|
| **GET**   | `/api/cars`    | `?filter={brand}`    | Get all cars (or with brand filter)  | `/api/cars?filter=Lamborghini`   |
| **POST**  | `/api/orders`  | `{car, name, phone}` | Create new booking                   | [See payload example](#)         |
| **GET**   | `/images/*`    | -                    | Access car images                    | `/images/3.jpg`  


// Request Payload
{
  "car": "Lamborghini",
  "name": "John Smith",
  "phone": "+971501234567"
}

// Success Response
{
  "success": true,
  "orderId": "ORD-789012",
  "message": "Booking confirmed! We'll contact you shortly."
}

<style>
  table {
    border-collapse: collapse;
    width: 100%;
    margin: 25px 0;
    box-shadow: 0 2px 15px rgba(0,0,0,0.1);
  }
  th {
    background-color: #2c3e50;
    color: white;
    text-transform: uppercase;
  }
  td, th {
    padding: 12px 15px;
    text-align: left;
    border-bottom: 1px solid #dddddd;
  }
  tr:hover {
    background-color: #f5f5f5;
  }
</style>

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