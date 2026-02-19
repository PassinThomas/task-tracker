BINARY_NAME=task
BUILD_DIR=bin

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/task

test:
	go test -v ./...

run:
	./$(BUILD_DIR)/$(BINARY_NAME) list

clean:
	rm -rf $(BUILD_DIR)

.PHONY: build run clean