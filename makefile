PROJECT_NAME=lamoda

build:
	docker-compose build

run:
	docker-compose up -d


up: build run

stop:
	docker stop $(PROJECT_NAME)-app-container $(PROJECT_NAME)-db-container

clean: stop
	docker rm $(PROJECT_NAME)-app-container $(PROJECT_NAME)-db-container
