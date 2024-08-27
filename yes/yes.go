package main

import (
	"fmt"
	"os"
)


type ErrTypes string
const ErrCommandNotExist = ErrTypes("the following command doesn't exist")
func (e ErrTypes) Error() string {
	return string(e)
}


func Yes(output string) {
	for {
		fmt.Println(os.Args[1])
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(ErrCommandNotExist)
		os.Exit(0)
	}
	Yes(os.Args[1])	
}