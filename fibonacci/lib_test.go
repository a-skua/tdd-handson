package fibonacci

import (
	"testing"
)

func TestFib(t *testing.T) {
	result := Fib(0)
	var want uint = 0
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}

	result = Fib(1)
	want = 1
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}

	result = Fib(2)
	want = 1
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}

	result = Fib(3)
	want = 2
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}

	result = Fib(4)
	want = 3
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}

	result = Fib(5)
	want = 5
	if result != want {
		t.Errorf("want=%v, got=%v", want, result)
	}

	// result = Fib(-1)
	// want = 0
	// if result != want {
	// 	t.Errorf("want=%v, got=%v", want, result)
	// }
}
