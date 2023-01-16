FROM golang:1.19-alpine

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY tests/go.mod ./
COPY tests/go.sum ./
RUN go mod download

COPY tests/. .
CMD ["go", "test"]