default:
	go run main.go

test:
	go test ./lexer

install:
	go install

clean:
	go clean

.PHONY: clean
