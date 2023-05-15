.PHONY: help
help: 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Compile the app and package in a docker image
.PHONY: docker-build
docker-build:
	docker build --tag docker-go-api -f Dockerfile.multistage .

# Push the Docker Image to DockerHub
# .PHONY: docker-build
# docker-push:
#	docker image push davidoram/kratos-selfservice-ui-go:latest

# Run the docker image
.PHONY: docker-run -l test-api
docker-run:
	docker run --detach --publish 3000:8080 docker-go-api

# List docker containers
.PHONY: docker-ls 
docker-ls:
	docker ps -a -q --filter ancestor=docker-go-api --format="{{.ID}}"

# Delete the Docker Containers
.PHONY: docker-rm-container
docker-rm-container:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor=docker-go-api --format="{{.ID}}"))

# Delete the Docker Images
.PHONY: docker-rm-image
docker-rm-image:
	docker image rm -f docker-go-api

# Build and run the Docker Container
up: docker-build docker-run

# Stop and Delete the Docker Containers and Images
.PHONY: down
down: 
# Check if there are running containers and stop them
ifneq ($(shell docker ps -q --filter ancestor=docker-go-api --format="{{.ID}}"),)
	@echo "Stop running containers: $(shell docker stop $$(docker ps -a -q --filter ancestor=docker-go-api --format="{{.ID}}"))"
else
	@echo "No running containers to stop"
endif
# Check if there are any containers and remove them
ifneq ($(shell docker ps -a -q --filter ancestor=docker-go-api --format="{{.ID}}"),)
	@echo "Remove all containers: $(shell docker rm $$(docker ps -a -q --filter ancestor=docker-go-api --format="{{.ID}}"))"
else
	@echo "No containers to remove"
endif
# Check if there are any images and remove them
ifneq ($(shell docker images -q docker-go-api),)
	@echo "Remove all images: $(shell docker image rm -f docker-go-api)"
else
	@echo "No images to remove"
endif



# test:
# 	go test ./...

# create_migration:
# 	migrate create -ext sql -dir migrations -seq create_url_table

# lint:
# 	golangci-lint run ./...


# mkdir -p $GOPATH/src/github.com/ory/kratos
# cd $GOPATH/src/github.com/ory/kratos
# git clone https://github.com/ory/kratos.git .
# git checkout v0.7.1-alpha.1
# docker build -t ory-kratos:v0.7.1-alpha.1 -f .docker/Dockerfile-build .

# mkdir -p $GOPATH/src/github.com/ory/kratos-selfservice-ui-node
# cd $GOPATH/src/github.com/ory/kratos-selfservice-ui-node
# git clone https://github.com/ory/kratos-selfservice-ui-node.git .
# git checkout v0.7.1-alpha.1
# docker build -t ory-kratos-selfservice-ui-node:v0.7.1-alpha.1 .