default:
	go build

test:
	go test ./lexer

install:
	go install

clean:
	go clean

.PHONY: clean
