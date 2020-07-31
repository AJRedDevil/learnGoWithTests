// Dependecny Injection
// - No need of framework
// - It does not overcomplicate design
// - It facilitates testing
// - It allows to write great, general-purpose functions.
package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

// Injeting a dependency
// - Test our code
//   If you can't test a function easily, it's usually because of
//   dependencies hard-wired into a function or global state. If
//   you have a global database connection pool for instance that
//   is used by some kind of service layer, it is likely going to
//   be difficult to test and they will be slow to run. DI will
//   motivate you to inject in a database dependency (via an interface)
//   which you can then mock out with something you can control in
//   your tests.
// - Separate our concerns
//   decoupling where the data goes from how to generate it. If you
//   ever feel like a method/function has too many responsibilities
//   (generating data and writing to a db? handling HTTP requests
//   and doing domain level logic?) DI is proably going to be the
//   tool you need
// - Allow our code to be re-used in different contexts
//   The first "new" context our code can be used in is inside tests.
//   But further on if someone wants to try something new with your
//   function they can inject their own dependencies.
