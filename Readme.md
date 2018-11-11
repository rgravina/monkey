# Monkey

This is an interpreter for Monkey, a programming language and interpreter from the book [Writing an Interpreter in Go](https://interpreterbook.com).

Here is a small example of Monkey.
```
let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
```
The code for this interpreter I have entered as I worked through the book, but I've also made small refactorings which helped me to understand it better. For the original source, please refer to the book.