package patternopt

import (
	"fmt"
	"strings"
)

func Pattern20(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%s%s%s\n", strings.Repeat("*", i), strings.Repeat(" ", n*2-i*2), strings.Repeat("*", i))
	}
	for i := n - 1; i > 0; i-- {
		fmt.Printf("%s%s%s\n", strings.Repeat("*", i), strings.Repeat(" ", n*2-i*2), strings.Repeat("*", i))
	}
}
