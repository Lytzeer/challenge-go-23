package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
    for , mot := range a {
        for , ch := range []rune(mot) {
            z01.PrintRune(ch)
        }
        z01.PrintRune('\n')
    }
}