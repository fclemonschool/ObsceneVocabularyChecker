package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error: am I the original error?")
	err = fmt.Errorf("error: no! I'm the original error! %w", err)

	// call the errors.Unwrap function here and pass err as the argument
	unwrappedErr := errors.Unwrap(err)

	fmt.Println(unwrappedErr)
	// print the unwrapped error here
}
