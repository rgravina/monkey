test:
	go test ./lexer ./parser

run:
	go run main.go

install:
	go install

clean:
	go clean

.PHONY: clean
