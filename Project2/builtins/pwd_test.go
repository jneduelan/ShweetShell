package builtins

import (
	"testing"
)

func TestGetworkDirectory(t *testing.T) {
	// Call GetworkDirectory and check for errors.
	if err := GetworkDirectory(); err != nil {
		t.Errorf("GetworkDirectory failed: %v", err)
	}
}


