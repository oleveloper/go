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

	r := rand.Intn(100)
	cnt := 1

	for {
		fmt.Println("Enter an integer value")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("Please enter an integer only")
		} else {
			if n > r {
				fmt.Println("The number entered is greater than the answer")
			} else if n < r{
				fmt.Println("The number entered is less than the answer")
			} else {
				fmt.Println("You got the numbers right! Congratulations! Number of attempts:", cnt)
				break
			}
			cnt++
		}
	}
}