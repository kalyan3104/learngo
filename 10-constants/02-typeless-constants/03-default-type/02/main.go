// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const min = 1

	max := 5 + min
	// above line equals to this:
	// max := int(5) + int(min)

	fmt.Printf("Type of max: %T\n", max)
}
