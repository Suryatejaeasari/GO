package patternopt

import (
	"fmt"
	"strings"
)

// Pattern9 prints a diamond pattern of stars.
func Pattern9(n int) {
	// Top half of the diamond
	for i := 1; i <= n; i++ {
		// Calculate spaces and stars
		spaces := n - i
		stars := i*2 - 1

		// Print the line
		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
	}

	// Bottom half of the diamond
	for i := n; i > 0; i-- {
		// Calculate spaces and stars
		spaces := n - i
		stars := i*2 - 1

		// Print the line
		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
	}
}
