# rterror

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/rterror)
[![Build Status](https://travis-ci.org/suzuki-shunsuke/rterror.svg?branch=master)](https://travis-ci.org/suzuki-shunsuke/rterror)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/rterror/master/LICENSE)

golang package to handle runtime error

## Install

```
$ go get github.com/suzuki-shunsuke/rterror
```

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/rterror"
)

func hello() (msg string, e error) { // declare named return value
	defer func() {
		p := recover()
		if p != nil {
			// assign RuntimeError to the surrounding function's return value
			e = rterror.NewRuntimeError(p, "")
		}
	}()
	var s *string
	fmt.Println(*s) // panic
	return "hello", nil
}

func main() {
	msg, err := hello()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
```

## Change Log

See [CHANGELOG.md](CHANGELOG.md).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)
