default:
	go run main.go

test:
	go test ./lexer ./parser

install:
	go install

clean:
	go clean

.PHONY: clean
