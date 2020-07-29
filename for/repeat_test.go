package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("b", 5)
	expected := "bbbbb"

	if repeated != expected {
		t.Errorf("expted %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("y", 3)
	fmt.Println(repeatedString)
	// Output: yyy
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 5)
	}
}
