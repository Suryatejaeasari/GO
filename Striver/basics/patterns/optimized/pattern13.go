package patternopt

import (
	"fmt"
	"strings"
)

func Pattern13(n int) {
	j := 1
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		for f := 1; f <= i; f++ {
			sb.WriteString(fmt.Sprintf("%d ", j))
			j++
		}
		fmt.Println(sb.String())
	}
}
