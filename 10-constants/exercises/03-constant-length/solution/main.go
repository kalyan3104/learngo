// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n",
		length, home)
}
