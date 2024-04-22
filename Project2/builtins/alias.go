package builtins

import (
    "errors"
    "strings"
)

// Aliases store command shortcuts
var Aliases = make(map[string]string)

// CreateAlias creates a new alias
func CreateAlias(args ...string) error {
    if len(args) != 1 {
        return errors.New("usage: alias name=command")
    }
    parts := strings.SplitN(args[0], "=", 2)
    if len(parts) != 2 {
        return errors.New("invalid alias format. Use name=command")
    }
    name := parts[0]
    command := parts[1]

    // Check if the command includes only space or if it is empty
    if strings.TrimSpace(command) == "" {
        return errors.New("command for alias cannot be empty")
    }

    Aliases[name] = command
    return nil
}

