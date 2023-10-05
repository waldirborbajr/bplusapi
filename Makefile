.PHONY: up
up:
	docker compose up -d

.PHONY: restart
restart:
	docker compose restart

.PHONY: logs
log:
	docker logs -f service-core

.PHONY: down
down:
	docker compose down

.PHONY: psql
psql:
	docker exec -it service-core-db psql -U postgres service-core

run-service-core-tests:
	cd service-core && go test ./...

generate-service-core-protoc:
	protoc --go_out=./service-core --go-grpc_out=./service-core ./service-core/protos/*.proto

# include .env
# export

# .PHONY: all
# all: build
# FORCE: ;

# .PHONY: build

# build: run-api run-app

# run-api:
# 	cd cmd/api && go mod tidy && go mod download && \
# 	CGO_ENABLED=0 go run github.com/thangchung/go-coffeeshop/cmd/product
# .PHONY: run-api

# run-app:
# 	cd cmd/app && go mod tidy && go mod download && \
# 	CGO_ENABLED=0 go run -tags migrate github.com/thangchung/go-coffeeshop/cmd/counter
# .PHONY: run-app


# test: generate-mocks
# 	go test ./...
# 	cd auth ; go test ./...
# 	cd feedbacks; go test ./...
# 	cd votes; go test ./...

# run-docker: build-auth-docker build-feedbacks-docker build-votes-docker
#     docker run -d -p 8081:8081 auth
#     docker run -d -p 8082:8082 feedbacks
#     docker run -d -p 8083:8083 votes

