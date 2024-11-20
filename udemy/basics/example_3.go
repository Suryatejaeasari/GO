package basics

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GuessingGame() {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(50)
	fmt.Println("Random Number: ", a)
	i := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your guess:")
	for true {
		v, err := i.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		v = strings.TrimSpace(v)
		iv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		if iv == a {
			fmt.Println("Correct Guess")
			break
		} else if iv < a {
			fmt.Println("Higher")
		} else {
			fmt.Println("Lower")
		}
	}

}
