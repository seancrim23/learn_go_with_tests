package iterate

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 8)
	fmt.Println(repeated)
	// Output: aaaaaaaa
}
