NAME := anavi-phat-storage

.DEFAULT_GOAL: $(NAME)

$(NAME): lint
	go build -v -o ${NAME} src/*.go

clean:
	go clean

test:
	go test -v ./src/*.go

lint:
	go vet ./src/*.go

run:
	go run src/*.go

.PHONY: clean test lint run
