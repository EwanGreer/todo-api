# TODO: add version

build:
	docker build -t app .

run: build
	docker run -it --rm --name app app

deploy: build
	@echo 'todo'
