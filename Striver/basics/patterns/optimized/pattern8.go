package patternopt

import (
	"fmt"
	"strings"
)

func Pattern8(n int) {
	for i := n; i > 0; i-- {
		st := i*2 - 1
		sp := n - i
		fmt.Println(strings.Repeat(" ", sp), strings.Repeat("*", st))
	}
}
