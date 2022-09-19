package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	n_tab := []string{}
	for i := 1; i < len(args); i++ {
		n_tab = append(n_tab, args[i])
	}
	for i := 0; i < len(n_tab); i++ {
		for j := 0; j < len(n_tab); j++ {
			min := n_tab[i]
			if n_tab[i] < n_tab[j] {
				fmt.Println(n_tab)
				n_tab[i] = n_tab[j]
				n_tab[j] = min
				fmt.Println(n_tab)
			}
		}
	}
	for i := 0; i < len(n_tab); i++ {
		for _, ch := range n_tab[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
