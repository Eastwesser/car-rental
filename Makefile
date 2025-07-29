.PHONY: full
full:
	docker-compose down
	docker-compose up --build