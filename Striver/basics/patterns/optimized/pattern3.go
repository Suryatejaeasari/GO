package patternopt

import (
	"fmt"
	"strings"
)

func Pattern3(n int) {
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		for j := 1; j <= i; j++ {
			sb.WriteString(fmt.Sprintf("%d", j))
		}
		fmt.Println(sb.String())
	}
}
