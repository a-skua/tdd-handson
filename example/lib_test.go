package example

import (
	"math/rand"
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

func TestNewUserID(t *testing.T) {
	// usage
	t.Log(NewUserID(rand.Int))

	result := NewUserID(func() int { return 1234567890 })
	want := "1234567890"
	if result != want {
		t.Errorf("result=%s, want=%s", result, want)
	}

	result = NewUserID(func() int { return 123456789 })
	want = "0123456789"
	if result != want {
		t.Errorf("result=%s, want=%s", result, want)
	}
}
