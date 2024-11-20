package patternopt

import (
	"fmt"
	"strings"
)

func Pattern19(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%s%s%s\n", strings.Repeat("*", n-i), strings.Repeat(" ", i*2), strings.Repeat("*", n-i))
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%s%s%s\n", strings.Repeat("*", i), strings.Repeat(" ", n*2-i*2), strings.Repeat("*", i))
	}
}
