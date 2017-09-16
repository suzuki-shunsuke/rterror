package rterror

import (
	"fmt"
	"log"
	"testing"

	"github.com/suzuki-shunsuke/rterror"
)

func TestDefaultError(t *testing.T) {
	err := rterror.NewRuntimeError(1, "")
	actual := err.Error()
	expected := "Runtime Error"
	if actual != expected {
		t.Errorf(
			"expected default Error(): %s, actual: %s",
			expected, actual)
	}
}

func TestStringError(t *testing.T) {
	expected := "hello"
	err := rterror.NewRuntimeError(expected, "")
	actual := err.Error()
	if actual != expected {
		t.Errorf(
			"expected Error(): %s, actual: %s",
			expected, actual)
	}
}

func TestParam(t *testing.T) {
	expected := 1
	err := rterror.NewRuntimeError(expected, "")
	actual := err.Param()
	if actual != expected {
		t.Errorf(
			"expected Param(): %s, actual: %s",
			expected, actual)
	}
}

func TestError(t *testing.T) {
	err := rterror.NewRuntimeError("", "")
	if _, ok := err.(error); !ok {
		t.Errorf("rterror.RuntimeError instance should be error")
	}
}

func ExampleRuntimeError() {
	msg, err := func() (msg string, e error) { // declare named return value
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
	}()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
