// Dependecny Injection
// - No need of framework
// - It does not overcomplicate design
// - It facilitates testing
// - It allows to write great, general-purpose functions.
package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
