// Handle runtime error.
package rterror

// Runtime Error interface
type RuntimeError interface {
	Param() interface{}
	Error() string
}

type runtimeErrorImpl struct {
	data struct {
		param interface{}
		msg   string
	}
}

// Get a value given to panic
func (e runtimeErrorImpl) Param() interface{} {
	return e.data.param
}

func (e runtimeErrorImpl) Error() string {
	return e.data.msg
}

// RuntimeError's constructor.
// `p` is a value given to panic.
// `msg` is a return value of Error().
// If `msg` is a empty string ("") and `p` holds a `string`, Error() returns the underlying value.
// If `msg` is a empty string ("") but `p` doesn't hold a string, `Error()` returns a string "Runtime Error".
func NewRuntimeError(p interface{}, msg string) RuntimeError {
	e := runtimeErrorImpl{}
	e.data.param = p
	if msg == "" {
		if s, ok := p.(string); ok {
			e.data.msg = s
		} else {
			e.data.msg = "Runtime Error"
		}
	}
	return e
}
