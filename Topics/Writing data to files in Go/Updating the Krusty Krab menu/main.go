package main

import (
	"fmt"
	"io"
	"strings"
)

var _ strings.Builder // do not remove this line!

func solve(file io.Writer) error {

	// write "Kelp Shake $2.00" to the 'file' variable below with the fmt.Fprintln() function

	_, err := fmt.Fprintln(file, "Kelp Shake $2.00")
	return err // do not delete this line!
}
