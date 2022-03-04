package main

// Do not remove any imports!
import (
	"errors"
	"fmt"
)

func brokenMessage() error {
	// create the main error here - you can use errors.New or fmt.Errorf
	err := errors.New("error: can't serve you an ice cream main cause of error: ice cream machine is broken")

	// create the wrapped error here using fmt.Errorf and the %w verb
	wrappedErr := fmt.Errorf("%w", err)

	return wrappedErr // Do not delete this line!
}
