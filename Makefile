.PHONY: cars
cars:
	docker-compose down
	docker-compose up --build