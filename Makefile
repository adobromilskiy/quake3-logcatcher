.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test
## test: runs tests
test:
	@go test -mod=vendor ./app/... -coverprofile cover.out

.PHONY: build
## build: builds application
build:
	@cd app && go build -v -mod=vendor

.PHONY: image
## image: build docker image
image:
	@docker build -t adobromilskiy/quake3-logcatcher .

.PHONY: run
## run: run docker container
run:
	@docker run -v /var/run/docker.sock:/run/docker.sock --rm adobromilskiy/quake3-logcatcher:latest --dbconn="mongodb://mongodb:27017" --container="quake3-server"