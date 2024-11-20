package basics

import (
	"fmt"
	"time"
)

func TimeOut() {
	now := time.Now()
	fmt.Println(now.Hour(), now.Minute(), now.Second())

}
