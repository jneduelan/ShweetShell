package builtins

import (
	"testing"
)

func TestRemoveAlias_ValidAlias(t *testing.T) {
	// Setup: Define an alias to be removed.
	Aliases["ls"] = "dir"

	// Test removing the alias.
	err := RemoveAlias("ls")
	if err != nil {
		t.Errorf("RemoveAlias failed: %v", err)
	}

	// Verify that the alias has been removed.
	if _, exists := Aliases["ls"]; exists {
		t.Errorf("Alias 'ls' was not removed as expected")
	}
}

func TestRemoveAlias_InvalidAlias(t *testing.T) {
	// Try to remove an alias that does not exist.
	err := RemoveAlias("nonexistent")
	if err == nil {
		t.Errorf("Expected error when trying to remove a nonexistent alias, but got none")
	}
}

