package main

import "fmt"

func main() {
	// Code Here
	// https://codeforces.com/problemset/problem/69/A

	var testCases int
	fmt.Scan(&testCases)

	type Vector struct {
		x, y, z int
	}

	netForce := Vector{}
	for i := testCases; i > 0; i-- {
		currentForce := Vector{}
		fmt.Scan(&currentForce.x, &currentForce.y, &currentForce.z)

		netForce.x += currentForce.x
		netForce.y += currentForce.y
		netForce.z += currentForce.z
	}

	if netForce.x == 0 && netForce.y == 0 && netForce.z == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
