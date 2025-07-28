.PHONY: backend
backend:
	docker-compose down
	docker-compose up --build backend

.PHONY: full
full:
	docker-compose down
	docker-compose up --build