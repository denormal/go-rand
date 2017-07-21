# go-rand

Package `go-rand` provides wrappers around
[crypto/rand](https://golang.org/pkg/crypto/rand/) for retrieving random values
of a range of basic types.

```go
import "github.com/denormal/go-rand"

// return a random string of 16 runes
str, err := rand.String(16)
if err != nil {
    panic(err)
}

// return a random 32 bit signed integer
value, err := rand.Int32()
if err != nil {
    panic(err)
}
```

For more information see `godoc github.com/denormal/go-rand`.

## Installation

`go-rand` can be installed using the standard Go approach:

```go
go get github.com/denormal/go-rand
```

## License

Copyright (c) 2017 Denormal Limited

[MIT License](LICENSE)
