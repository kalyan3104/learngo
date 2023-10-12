// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"path"
)

func main() {
	_, file := path.Split("css/main.css")

	// or this:
	// dir, file := path.Split("css/main.css")

	fmt.Println("file:", file)
}
