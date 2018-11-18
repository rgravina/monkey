test:
	go test ./lexer ./parser ./ast ./object

run:
	go run main.go

install:
	go install

clean:
	go clean

.PHONY: clean
