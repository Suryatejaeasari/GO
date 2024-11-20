package patternopt

import (
	"fmt"
	"strings"
)

func Pattern21(n int) {
	for i := 1; i <= n; i++ {
		if i == 1 || i == n {
			fmt.Println(strings.Repeat("*", n))
		} else {
			fmt.Printf("*%s*\n", strings.Repeat(" ", n-2))
		}
	}
}
