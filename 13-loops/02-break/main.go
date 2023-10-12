// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}

		sum += i

		fmt.Println(i, "-->", sum)

		i++
	}

	fmt.Println(sum)
}
