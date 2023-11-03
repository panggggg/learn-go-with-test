package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat P", func(t *testing.T) {
		repeated := Repeat("P", 5)
		expected := "PPPPP"

		if repeated != expected {
			t.Errorf("expected %q bot got %q", expected, repeated)
		}
	})
	t.Run("repeat P for 7 times", func(t *testing.T) {
		repeated := Repeat("P", 7)
		expected := "PPPPPPP"

		if repeated != expected {
			t.Errorf("expected %q bot got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
	// When the benchmark code is executed, it runs b.N times and measures how long it takes.
	// the framework will determine what is a "good" value for that to let you have some decent results.
	// go test -bench=.
}

func ExampleRepeat() {
	repeated := Repeat("P", 4)
	fmt.Println(repeated)
	// Output: PPPP
}
