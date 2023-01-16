FROM golang:1.19-alpine

WORKDIR /opt/integration_tests

COPY ./tests/go.mod .
COPY ./tests/go.sum .
RUN go mod download

COPY ./tests/ .