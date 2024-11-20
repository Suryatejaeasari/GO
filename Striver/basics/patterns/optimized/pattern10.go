package patternopt

import (
	"fmt"
	"strings"
)

func Pattern10(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
	for i := n - 1; i > 0; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}
