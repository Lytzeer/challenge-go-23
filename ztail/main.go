package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	n := Atoi(args[0])
	tab := []string{}
	args = args[1:]
	word := ""
	for i := 1; i < len(args); i++ {
		tab = append(tab, args[i])
	}
	for i := 1; i < len(tab); i++ {
		word = ""
		for j := (len(tab[i]) - 1) - n; j < len(tab[i]); j++ {
			word += string(tab[i][j])
		}
		fmt.Print(word, "\n")
	}
}

func Atoi(s string) int {
	atoi := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			atoi = atoi*10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return atoi
}
