package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func print_rune(message string) {
	for _, i := range message {
		z01.PrintRune(i)
	}
}

func main() {
	if len(os.Args) == 1 {
		content, _ := ioutil.ReadAll(os.Stdin)
		print_rune(string(content))
	} else {
		args := os.Args[1:]

		for i := range args {
			content, err := ioutil.ReadFile(args[i])
			if err == nil {
				print_rune(string(content))
			} else {
				print_rune("ERROR: open " + args[i] + ": no such file or directory")
				print_rune("\n")
				os.Exit(1)
			}
		}
	}
}
