package patternopt

import (
	"fmt"
	"strings"
)

func Pattern7(n int) {
	for i := 1; i <= n; i++ {
		st := i*2 - 1
		sp := n - i
		fmt.Println(strings.Repeat(" ", sp), strings.Repeat("*", st))
	}
}
