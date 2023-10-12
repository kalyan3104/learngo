// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	var counter int

	counter++
	counter--

	counter += 5
	counter *= 10
	counter /= 5

	fmt.Println(counter)
}
