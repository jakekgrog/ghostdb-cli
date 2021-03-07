GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_GET=$(GO_CMD) get
NAME=ghostdbcli
DIR=./cmd/ghostdbcli

all: build
build:
	$(GO_BUILD) -o $(NAME) -v $(DIR)
install:
	$(GO_BUILD) -o /bin/$(NAME) -v $(DIR)
clean:
	$(GO_CLEAN) -v ./...
	rm -f $(NAME)
run:
	./$(NAME)