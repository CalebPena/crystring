build:
	go build -o bin/dont_forget

run: build
	./bin/dont_forget

test:
	go test -v ./...
