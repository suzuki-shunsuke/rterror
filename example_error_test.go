package rterror_test

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

func Example() {
	msg, err := hello()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
