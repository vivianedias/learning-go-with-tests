package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' five times", func (t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repeat 'b' 2 times", func (t *testing.T) {
		repeated := Repeat("b", 2)
		expected := "bb"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("c", 10)
	fmt.Println(repeat)
	// Output: cccccccccc
}