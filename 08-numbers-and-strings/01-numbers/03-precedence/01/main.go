// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	fmt.Println(
		2+2*4/2,
		2+((2*4)/2), // same as above
	)

	fmt.Println(
		1+4-2,
		(1+4)-2, // same as above
	)

	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2), // same as above
	)
}
