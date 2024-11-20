package patternopt

import (
	"fmt"
	"strings"
)

func Pattern4(n int) {
	// for i := 1; i <= n; i++ {
	// 	var sb strings.Builder
	// 	for j := 1; j <= i; j++ {
	// 		sb.WriteString(fmt.Sprintf("%d", i))
	// 	}
	// 	fmt.Println(sb.String())
	// }
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat(fmt.Sprintf("%d", i), i))
	}
}
