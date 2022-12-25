
<div align="center">
    <img src="https://i.imgur.com/fkHEaWY.png">

# DevScript

## Programming Language

</div>

*DevScript Logo by <a href="https://dribbble.com/scss04">Sanatan Choudhary</a>*

## Features

- [x] Comment
- [x] Literal types: int, bool, string
- [x] Expression evaluation
- [x] Variable Declaration and initialization
- [x] Higher level function
- [x] If-Else Expression
- [x] Builtin len, print & println functions
- [x] REPL
- [x] Run `.ds` File

## Build

Run `make build` to build the executable.
If Golang is installed on your system, build using `go build` command.

## Hands On with the language

Create a new `.ds` file and write the following code.

```ds
var a = 10;
var b = 5;

var add = func(x, y) {
    return x + y;
}

var sub = func(x, y) {
    return x - y;
}

var operation = func(x, y, fn) {
    return fn(x,y);
}

var result = operation(a, b, add);

func addition(x, y) {
    return x + y;
}

if (result == addition(a, b)) {
    print("The Result is:");
    println(result);
} else {
    print("Wrong Result");
}
```

## Contributions and suggestions

Refer `CONTRIBUTING.md` file for more information about the contribution and suggestions.
