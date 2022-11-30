
$(NAME)

all: build

$(NAME): build

build:
	go build .

run:
	go run .

.PHONY: all 
