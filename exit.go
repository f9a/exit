package exit

import (
	"fmt"
	"os"
)

// Code represents the programm exit code
type Code int

// Catch handles the graceful exit of the program.
// If os.Exit is called directly in any other function the deferred functions are NOT executed.
// Therefore Catch must be called within the first defer statement in the main function.
// If then a function like to exit the program panic(ExitCode(int)) must be called.
func Catch() {
	if e := recover(); e != nil {
		if code, ok := e.(Code); ok == true {
			os.Exit(int(code))
		}
		panic(e) // not an Exit, bubble up
	}
}

// With panic with given code
func With(code int) {
	panic(Code(code))
}

// OnErrf exits with error-code=1 when err is not nil and print error to stdout.
func OnErrf(err error, format string, args ...interface{}) {
	if err != nil {
		msg := fmt.Sprintf(format, args...)
		fmt.Printf("%s: %v\n", msg, err)
		With(1)
	}
}

// OnErr exits with error-code=1 when err is not nil and print error to stdout.
func OnErr(err error) {
	if err != nil {
		fmt.Println(err)
		With(1)
	}
}

// WithErr exits with error-code=1 and print error to stdout.
func WithErr(err error) {
	fmt.Println(err)
	With(1)
}

// WithErrf exits with error-code=1 and printf the given message.
func WithErrf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	With(1)
}
