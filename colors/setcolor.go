package colors

import "fmt"

func Color(c, text string) {
	colorReset := "\033[97m"
	fmt.Print(string(c), text, "\n", string(colorReset))
}

var Colors = []string{
	"\033[31m", //cyan 0
	"\033[31m", //red 1
	"\033[33m", //yellow 2
	"\033[32m", //green 3
	"\033[34m", //purple 4
} 