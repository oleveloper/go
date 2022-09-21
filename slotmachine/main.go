package main

import (
	"fmt"
	"time"
	"math/rand"
	"bufio"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	money := 500

	for {
		r := rand.Intn(6)
		fmt.Println("Please enter a number between 1 and 5")
		n, err := InputIntValue()
		if err != nil || !(n > 0 && n < 6) {
			
		} else {
			if n == r {
				money += 500
				fmt.Println("Congratulations! You got 500 won! money:", money)
			} else {
				money -= 100
				fmt.Println("That's too bad. 100 won will be deducted. money:", money)
			}

			if money >= 5000 || money <= 0 {
				break
			}
		}
	}
}