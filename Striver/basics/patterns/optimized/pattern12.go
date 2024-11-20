package patternopt

import (
	"fmt"
	"strings"
)

func Pattern12(n int) {
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		for j := 1; j <= i; j++ {
			sb.WriteString(fmt.Sprintf("%d", j))
		}
		sb.WriteString(strings.Repeat(" ", n*2-i*2))
		for j := i; j > 0; j-- {
			sb.WriteString(fmt.Sprintf("%d", j))
		}
		fmt.Println(sb.String())
	}
}
