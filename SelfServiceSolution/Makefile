# Compile the App and package in a Docker Image

# Push the Docker Image to DockerHub


mkdir -p $GOPATH/src/github.com/ory/kratos
cd $GOPATH/src/github.com/ory/kratos
git clone https://github.com/ory/kratos.git .
git checkout v0.7.1-alpha.1
docker build -t ory-kratos:v0.7.1-alpha.1 -f .docker/Dockerfile-build .

mkdir -p $GOPATH/src/github.com/ory/kratos-selfservice-ui-node
cd $GOPATH/src/github.com/ory/kratos-selfservice-ui-node
git clone https://github.com/ory/kratos-selfservice-ui-node.git .
git checkout v0.7.1-alpha.1
docker build -t ory-kratos-selfservice-ui-node:v0.7.1-alpha.1 .