.PHONY: help
help: 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


### Build Go Solution ###
#########################
.PHONY: go
go:
	go run .

##############################################
### Docker Compose commands ###

# Run docker compose
.PHONY: up
up:
ifdef build
	@echo "Force build images and start docker compose"
	docker-compose up --build --force-recreate web-api
else
	@echo "Start docker compose"
	docker compose up
endif

# Stop docker-compose and remove images
down: remove-all-containers remove-web-api-image

# Stop docker-compose
down-c: remove-all-containers

# Stop and Delete the Docker Containers and Images
.PHONY: remove-all-containers
remove-all-containers:
	@echo "Remove all containers"
	docker compose down

# Stop and Delete the Docker Containers and Images
.PHONY: remove-web-api-image
remove-web-api-image:
	@echo "Remove web-api image"
	docker image rm -f sol-web-api

list: docker-ls-images docker-ls-containers

##############################################
### Docker file, image, container commands ###

# Compile app and package in a docker image
.PHONY: docker-build-image
docker-build-image:
	docker build --tag $(name) -f Dockerfile .

# Run the docker image
.PHONY: docker-run-image
docker-run-image:
	docker run --detach --publish 3000:8080 $(name)

# List docker images
.PHONY: docker-ls-images
docker-ls-images:
	@echo "LIST IMAGES"
	docker image ls $(name)

# Delete docker images
.PHONY: docker-rm-images
docker-rm-images:
	docker image rm -f $(name)

# List docker containers
.PHONY: docker-ls-containers
docker-ls-containers:
	@echo "LIST CONTAINERS"
	docker ps -a

# Delete the docker containers
.PHONY: docker-rm-containers
docker-rm-containers:
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor=$(name) --format="{{.ID}}"))

# Build the docker image and run in container
up-image: docker-build-image docker-run-image

# Stop and delete the docker containers and images
.PHONY: down-image
down-image: 
# Check if there are running containers and stop them
ifneq ($(shell docker ps -q --filter ancestor=$(name) --format="{{.ID}}"),)
	@echo "Stop running $(name) containers: $(shell docker stop $$(docker ps -a -q --filter ancestor=$(name) --format="{{.ID}}"))"
else
	@echo "No running $(name) containers to stop"
endif
# Check if there are any containers and remove them
ifneq ($(shell docker ps -a -q --filter ancestor=$(name) --format="{{.ID}}"),)
	@echo "Remove all $(name) containers: $(shell docker rm $$(docker ps -a -q --filter ancestor=$(name) --format="{{.ID}}"))"
else
	@echo "No $(name) containers to remove"
endif
# Check if there are any images and remove them
ifneq ($(shell docker images -q $(name)),)
	@echo "Remove all $(remove) images: $(shell docker image rm -f $(name))"
else
	@echo "No $(name) images to remove"
endif
# Ask if we want to prune any dangling images
	docker image prune




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