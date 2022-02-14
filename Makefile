.PHONY: phony
phony-goal: ; @echo $@

build: validate
	docker-compose -f containerd/docker-compose.yml up --build --remove-orphans --force-recreate --detach

validate: format vet lint test

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -covermode count -coverprofile coverage.out ./...

coverage: test
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

run-migrate-up:
	go run . migrate up

run-migrate-down:
	go run . migrate down

run-migrate-drop:
	go run . migrate drop

run-serve:
	go run . serve

run-test:
	go run . test

clean:
	rm -f coverage.* main.bin

compile:
	go build -a -o main.bin .

generate-stubs:
	protoc --go_out=. --go-grpc_out=. ./api/jukebox-app.proto

generate-mocks:
	go generate ./src/tests/mock

env-setup:
	docker-compose -f containerd/docker-compose.yml up --build --remove-orphans --force-recreate --detach jukebox-mysql
	sleep 2
	docker-compose -f containerd/docker-compose.yml up --build --remove-orphans --force-recreate --detach jukebox-redis
	docker-compose -f containerd/docker-compose.yml up --build --remove-orphans --force-recreate --detach jukebox-nats
	sleep 2
	go run . migrate up

prepare:
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go mod download
	go mod tidy
