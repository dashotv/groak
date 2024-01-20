include .env
export $(shell sed 's/=.*//' .env)

all: test

test:
	go test -v ./...

build:
	go build

install: build
	go install

docker:
	docker build -t groak .

docker-run:
	docker run --rm --name groak-test -p 19000:9000 groak

dotenv:
	npx dotenv-vault local build

.PHONY: test build install docker docker-run
