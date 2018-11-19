test:
	go test ./lexer ./parser ./ast ./evaluator

run:
	go run main.go

install:
	go install

clean:
	go clean

.PHONY: clean
