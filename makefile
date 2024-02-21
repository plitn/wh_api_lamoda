PROJECT_NAME=lamoda

build:
	docker-compose build

run:
	docker-compose up -d


up: build run

stop:
	docker stop wh_api_lamoda_golang_app_1
	docker stop wh_api_lamoda_db_1

clean: stop
	docker rm wh_api_lamoda_golang_app_1
	docker rm wh_api_lamoda_db_1
