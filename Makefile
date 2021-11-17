run:
	@echo "Executing run target..."
	go run main.go

build:
	@echo "Executing build target: installing and running..."
	go build -o bin/
	./bin/weblog

test:
	@echo "Executing test target..."
	go test -v -cover ./...