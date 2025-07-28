.PHONY: run-backend
run-backend:
	docker-compose down
	docker-compose up --build backend

.PHONY: run-full
run-full:
	docker-compose down
	docker-compose up --build