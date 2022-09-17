package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	fmt.Println(n)
}