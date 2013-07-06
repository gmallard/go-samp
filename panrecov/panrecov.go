package main

import (
	"errors"
	"fmt"
	"runtime"
)

// A blatant copy from and thanks to:
// http://www.goinggo.net/2013/06/understanding-defer-panic-and-recover.html

func main() {
	var err error

	err = TestFinal()

	if err != nil {
		fmt.Printf("Main Error: %v\n", err)
	}
}

func _CatchPanic(err *error, functionName string) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", functionName, r)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))

		if err != nil {

			*err = errors.New(fmt.Sprintf("%v", r))
		}
	} else if err != nil && *err != nil {

		fmt.Printf("%s : ERROR : %v\n", functionName, *err)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		fmt.Printf("%s : Stack Trace : %s", functionName, string(buf))
	}
}

func MimicError(key string) error {
	return errors.New(fmt.Sprintf("Mimic Error : %s", key))
}

func TestFinal() (err error) {
	defer _CatchPanic(&err, "TestFinal")

	fmt.Printf("Start Test\n")

	err = MimicError("1")

	panic("Mimic Panic")

	fmt.Printf("End Test\n")

	return err
}
