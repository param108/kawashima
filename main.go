package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	PLUS = 0
	MINUS = 1
	MULT = 2
	NumSums = 100
)

func main() {

	rand.Seed(time.Now().UnixNano())
	OpMap := map[int]string {
		PLUS: "+",
		MINUS: "-",
		MULT : "*",
	}

	onlySum := false
	for i := 0; i < NumSums; i++ {
		x := rand.Intn(20)
		y := rand.Intn(20)

		if x > 12 || y > 12 {
			onlySum = true
		}

		if y > x {
			x,y = y,x
		}

		ch := 0
		if onlySum {
			// only PlUS and MINUS allowed
			ch = rand.Intn(2)

		} else {
			// PLUS MINUS and MULT allowed
			ch = rand.Intn(3)
		}

		onlySum = false

		fmt.Printf("%2d %s %2d = _____", x, OpMap[ch], y)
		if i%3 != 2 {
			fmt.Printf("\t\t\t")
		} else {
			if i != NumSums - 1 {
				fmt.Print("\n\n")
			}
		}
	}

	if (NumSums - 1) %3 != 2 {
		fmt.Print("\n\n")
	}
}
