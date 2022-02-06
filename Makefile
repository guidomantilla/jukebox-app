.PHONY: phony
phony-goal: ; @echo $@

build: validate compile

validate: format vet lint test

generate:
	protoc --go_out=. ./api/jukebox-app.proto

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -covermode count -coverprofile coverage.out ./...

coverage-local: test
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


compile:
	go build -a -o /main .

env-setup:
	docker compose up -d jukebox-mysql

prepare:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go mod download
	go mod tidy
