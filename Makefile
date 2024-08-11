APP_NAME = todo-app
SRC = cmd/main.go

# build:
# 	go build -o $(APP_NAME) $(SRC)

# run: build
# 	./$(APP_NAME)

run:
	go run $(SRC)

clean:
	rm -f $(APP_NAME)

.PHONY: all run clean

all: run