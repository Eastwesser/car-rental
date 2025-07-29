.PHONY: backend
backend:
	docker-compose down
	docker-compose up --build backend

.PHONY: frontend
backend:
	docker-compose down
	docker-compose up --build frontend

.PHONY: full
full:
	docker-compose down
	docker-compose up --build