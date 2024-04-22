package builtins

import (
//	"errors"
	"fmt"
	"os"
)

// Define error messages related to argument counts to ensure they are consistent and maintainable.


// Renamefile handles renaming or moving files and directories. Expects two Arguments.
// args[0]: the current name or path of the file/directory
// args[1]: the new name or path to which the file/directory should be moved/renamed
func Renamefile(args ...string) error {
	// Check if the number of arguments is exactly 2
	if len(args) != 2 {
		return fmt.Errorf("%w", ErrInvalidArgCount)
	}

	// Attempt to rename/move the file with os.Rename
	err := os.Rename(args[0], args[1])
	if err != nil {
		
		return fmt.Errorf("error renaming/moving file: %w", err)
	}

	// Return nil if the operation was successful
	return nil
}

