package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber() {
	rand.Seed(time.Now().Unix())
	fmt.Printf("Number: %d", rand.Intn(50)+1)
}
