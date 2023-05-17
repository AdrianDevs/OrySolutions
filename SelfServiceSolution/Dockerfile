# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

# Set the below directory and destination for COPY and the working directory
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY public ./public
COPY templates ./templates

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /web-api

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /web-api /web-api
COPY --from=build-stage /app/public /public
COPY --from=build-stage /app/templates /templates

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
# If we wanted to expose port `8080` inside the container to port `3000` outside 
# the container, we would pass `3000:8080` to the `--publish` flag.
# EXPOSE 3000:3000

USER nonroot:nonroot

ENTRYPOINT ["/web-api"]