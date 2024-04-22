package builtins

import (
    "fmt"
)

// RemoveAlias removes an existing alias
func RemoveAlias(args ...string) error {
    if len(args) != 1 {
        return fmt.Errorf("usage: unalias name")
    }
    name := args[0]

    if _, exists := Aliases[name]; !exists {
        return fmt.Errorf("alias '%s' does not exist", name)
    }

    delete(Aliases, name)
    return nil
}


