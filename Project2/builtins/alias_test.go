package builtins

import (
	"testing"
)

func TestCreateAlias(t *testing.T) {
	err := CreateAlias("ls=dir")
	if err != nil {
		t.Errorf("CreateAlias failed: %v", err)
	}
	if Aliases["ls"] != "dir" {
		t.Errorf("Alias was not set correctly")
	}
}

func TestRemoveAlias(t *testing.T) {
	Aliases["ls"] = "dir"
	err := RemoveAlias("ls")
	if err != nil {
		t.Errorf("RemoveAlias failed: %v", err)
	}
	if _, exists := Aliases["ls"]; exists {
		t.Errorf("Alias was not removed")
	}
}

