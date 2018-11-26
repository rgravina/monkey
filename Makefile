test:
	go test ./lexer ./parser ./ast ./object ./evaluator

run:
	go run main.go

install:
	go install

clean:
	go clean

.PHONY: clean
