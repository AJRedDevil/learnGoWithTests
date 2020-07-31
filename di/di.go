// Dependecny Injection
// - No need of framework
// - It does not overcomplicate design
// - It facilitates testing
// - It allows to write great, general-purpose functions.
package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "AJ")
}
