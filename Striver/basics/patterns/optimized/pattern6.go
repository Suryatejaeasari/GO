package patternopt

import (
	"fmt"
	"strings"
)

func Pattern6(n int) {
	for i := n; i > 0; i-- {
		var sb strings.Builder
		for j := 1; j <= i; j++ {
			sb.WriteString(fmt.Sprintf("%d", j))
		}
		fmt.Println(sb.String())
	}
}
