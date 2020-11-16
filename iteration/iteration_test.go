package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("t", 10)
	expected := "tttttttttt"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkReport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	// Output "aaaaaaaaaa"
}
