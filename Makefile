build:
	go build -o bin/cdn cmd/cdn-fiber/main.go

run: build
	./bin/cdn

test:
	go test -v ./... -count=1