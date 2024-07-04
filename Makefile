build:
	go build -o bin/api cmd/api/main.go

run: build
	./bin/api

test-api:
	go test -v ./tests/... -count=1

