package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("File name missing")
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		data, err := ioutil.ReadFile(args[1])
		if err == nil {
			fmt.Printf(string(data))
		}
	}
}
