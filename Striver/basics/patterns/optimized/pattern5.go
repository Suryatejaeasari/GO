package patternopt

import (
	"fmt"
	"strings"
)

func Pattern5(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(strings.Repeat("* ", i))
	}
}
