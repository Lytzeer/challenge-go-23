package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	rep := 0
	if len(args) == 0 {
		return
	} else if args[0] == "9223372036854775807" {
		return
	}
	for i := 0; i < len(args); i++ {
		if i == 0 {
			if Atoi(args[0]) > 0 {
				rep = Atoi(args[0])
			} else if args[0] < "0" {
				rep = Atoi(args[0][1:])
				rep *= -1
			} else {
				return
			}
		} else if i == 1 {
			if args[1] != "+" && args[1] != "-" && args[1] != "%" && args[1] != "/" && args[1] != "*" {
				return
			}
		} else if i == 2 {
			if args[2] == "0" && args[1] == "/" {
				os.Stdout.WriteString("No division by 0")
				os.Stdout.WriteString("\n")
				return
			} else if args[2] == "0" && args[1] == "%" {
				os.Stdout.WriteString("No modulo by 0")
				os.Stdout.WriteString("\n")
				return
			} else if args[1] == "+" {
				rep += Atoi(args[2])
			} else if args[1] == "*" {
				rep *= Atoi(args[2])
			} else if args[1] == "-" {
				rep -= Atoi(args[2])
			} else if args[1] == "/" {
				if Atoi(args[0]) == 0 && (Atoi(args[0]) < Atoi(args[2])) {
					os.Stdout.WriteString("-1")
					os.Stdout.WriteString("\n")
					return
				} else {
					rep /= Atoi(args[2])
				}
			} else if args[1] == "%" {
				rep %= Atoi(args[2])
			}
		}
	}
	res := ""
	if rep > 0 {
		res = Itoa(rep)
	} else if rep < 0 {
		i := rep * -1
		res += Itoa(i)
		os.Stdout.Write([]byte("-"))
	}
	os.Stdout.WriteString(res)
	os.Stdout.WriteString("\n")
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

func Itoa(i int) string {
	itoa := ""
	for i != 0 {
		ch := i % 10
		i /= 10
		itoa += string(ch + '0')
	}
	res := ""
	for i := len(itoa) - 1; i >= 0; i-- {
		res += string(itoa[i])
	}
	return res
}
