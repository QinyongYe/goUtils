package main

import (
	. "github.com/cyfdecyf/color"
	"fmt"
)

func main() {
	colorF := []func(string) string{
		Red, Green, Yellow, Blue, Magenta, Cyan,
	}

	for _, f := range colorF {
		fmt.Println(f("sdfsdgsdg"))
	}
}
