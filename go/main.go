package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	var max_ = int64(math.Ceil(math.Sqrt(float64(n))))
	var i int64
	for i = 2; i <= max_; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var args = os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: %s <number>\n", args[0])
		os.Exit(1)
	}
	var n, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Fatalf("fail to parse number: %v", err)
	}
	run(n)
}

func run(n int64) {
	var count = 0
	startTime := time.Now()
	var i int64

	for i = 0; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}

	var elapse = time.Now().Sub(startTime)
	fmt.Printf("go, %d, %d, %v\n", n, count, elapse.Milliseconds())
}
