# TODO: add version

CONTAINER_NAME := app

build:
	docker build -t app .

run: build
	docker run -it --rm --name app ${CONTAINER_NAME}

up: 
	docker compose up --build -d

down:
	docker compose down

logs:
	docker logs -f ${CONTAINER_NAME}


