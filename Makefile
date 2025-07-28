docker-compose down
docker-compose up --build

Publish:
docker build -t eastwesser/car-rental-backend ./backend
docker build -t eastwesser/car-rental-frontend ./frontend

After - change in services:
services:
  backend:
    image: eastwesser/car-rental-backend
    # ...
  
  frontend:
    image: eastwesser/car-rental-frontend
    # ...