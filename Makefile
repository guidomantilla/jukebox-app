.PHONY: phony
phony-goal: ; @echo $@

build: validate
	docker compose -f docker/docker-compose.yml up --detach

build-portainer-stack:
	docker compose -f docker/portainer-stack-compose.yml up --detach

build-jaeger-stack:
	docker compose -f docker/jaeger-stack-compose.yml up --detach

build-sonarqube-stack:
	docker compose -f docker/sonarqube-stack-compose.yml up --detach

build-kafka-stack:
	docker compose -f docker/kafka-stack-compose.yml up --detach

build-mysql-stack:
	docker compose -f docker/mysql-stack-compose.yml up --detach

build-postgres-stack:
	docker compose -f docker/postgres-stack-compose.yml up --detach

validate: sort-import format vet lint sonarqube

generate:
	go generate ./...

sort-import:
	goimportssort -w ./..

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -covermode count -coverprofile coverage.out.tmp ./pkg/... ./internal/...
	cat coverage.out.tmp | grep -v "Mock" > coverage.out

coverage: test
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

sonarqube: coverage
	sonar-scanner

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

update-dependencies:
	go get -u ./...
	go get -t -u ./...
	go mod tidy

env-setup:
	docker compose -f containerd/docker-compose.yml up --build --remove-orphans --force-recreate --detach jukebox-mysql jukebox-redis
	sleep 10
	go run . migrate up

prepare:
	go install github.com/AanZee/goimportssort@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/ktr0731/evans@latest
	go mod download
	go mod tidy
