package builtins

import "testing"

func TestGettime(t *testing.T) {
	// Testing Gettime which only prints time. You could capture stdout if necessary.
	if err := Gettime(); err != nil {
		t.Errorf("Gettime failed: %v", err)
	}
}

