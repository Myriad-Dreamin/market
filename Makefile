all:
	go build .

build-docker:
	docker build --tag minimum-market/backend:alpine .

run-docker:
	docker run -id -p 23336:23336 --name backend minimum-market/backend:alpine


.PHONY: all build-docker



