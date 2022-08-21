package fibonacci

import (
	"testing"
)

func TestFib(t *testing.T) {
	result := Fib(0)
	want := 0
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}
}
