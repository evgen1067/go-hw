build:
	go build -o calendar cmd/calendar/main.go

run:
	go run cmd/calendar/main.go

lint:
	golangci-lint run ./...

test:
	go test -v -count=1 -race -timeout=30s ./...

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        api/EventService.proto

integration-tests:
	set -e ;\
    	docker compose -f ./deployments/docker-compose.test.yaml up --build -d ;\
    	test_status_code=0 ;\
    	docker compose -f ./deployments/docker-compose.test.yaml run integration_tests go test || test_status_code=$$? ;\
    	docker compose -f ./deployments/docker-compose.test.yaml down ;\
    	exit $$test_status_code ;

test-cleanup:
	docker compose -f ./deployments/docker-compose.test.yaml down \
        --rmi local \
		--volumes \
		--remove-orphans \
		--timeout 60; \
  	docker compose rm -f

up:
	docker compose -f ./deployments/docker-compose.yaml up

rebuild:
	docker compose -f ./deployments/docker-compose.yaml up --build

down:
	docker compose -f ./deployments/docker-compose.yaml down --remove-orphans

create_db:
	docker exec -it otus_go_postgres createdb --username=go_user --owner=go_user events_db

migrate_up:
	docker compose -f ./deployments/docker-compose.yaml run migrate -path /migrations -database "postgres://go_user:go_password@database:5432/events_db?sslmode=disable" up

migrate_down:
	docker compose -f ./deployments/docker-compose.yaml run migrate -path /migrations -database "postgres://go_user:go_password@database:5432/events_db?sslmode=disable" down

.PHONY:


