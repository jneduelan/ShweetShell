package builtins

import (
	"fmt"
	"time"
)

// Gettime prints the current system time.
// Doesn't require any agurments, any provided arguments are ignored.
func Gettime(args ...string) error {
	// Obtain the current time using time.Now()
	mytime := time.Now()

	// Format the time to a more readable form 
	fmt.Println("Current time:", mytime.Format("2006-01-02 15:04:05"))

	return nil
}

