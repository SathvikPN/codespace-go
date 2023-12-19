package main

import (
	"fmt"
	"math/rand"
)

const (
	LOOP_MAX = 100
	LOOP_MIN = 1
)

var (
	GENERATOR = func() int {
		return rand.Intn(100)
	}
)

func main() {

	// Print Test Case Structure Here

	LOOP := rand.Intn(LOOP_MAX-LOOP_MIN) + LOOP_MIN
	fmt.Println(LOOP)

	for i := LOOP; i > 0; i-- {

		fmt.Printf("%v %v %v\n",
			GENERATOR(), GENERATOR(), GENERATOR())
	}
}
