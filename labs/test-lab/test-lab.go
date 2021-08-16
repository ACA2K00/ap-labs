package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1:]
	if len(os.Args) == 1 {
		fmt.Printf("No name submitted, submit your name please... \n")
	}else{
		fmt.Printf("Welcome %v to the jungle \n", name)
	}
}
