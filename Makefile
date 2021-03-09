GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_GET=$(GO_CMD) get
NAME=gdb
DIR=./cmd

all: build
build:
	$(GO_BUILD) -o $(NAME) -v $(DIR)
buildWin:
	$(GO_BUILD) -o $(NAME).exe -v $(DIR)
install:
	$(GO_BUILD) -o /bin/$(NAME) -v $(DIR)
clean:
	$(GO_CLEAN) -v ./...
	rm -f $(NAME)
run:
	./$(NAME)