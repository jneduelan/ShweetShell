package builtins

import (
	"fmt"
	"os"
)

// GetworkDirectory prints the current working directory.
// Doesn't require arguments, any provided arguments are ignored.
func GetworkDirectory(args ...string) error {
	// Get current working directory using os.Getwd()
	mydir, err := os.Getwd()
	if err != nil {
	
		return fmt.Errorf("error getting working directory: %w", err)
	}
	
	// Print current working directory
	fmt.Println(mydir)
	return nil
}
