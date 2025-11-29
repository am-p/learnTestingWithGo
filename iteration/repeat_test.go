package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	want := Repeat("a", 5)
	got := "aaaaa"

	if want != got {
		t.Errorf("want '%q' but got '%q'", got, want)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: 'bbb'
}
