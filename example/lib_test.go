package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	if result != 7 {
		t.Fatalf("want=%d, got=%d", 7, result)
	}
}

func TestAdd3(t *testing.T) {
	result := Add3(3, 4, 5)
	if result != 12 {
		t.Fatalf("want=%d, got=%d", 12, result)
	}
}
