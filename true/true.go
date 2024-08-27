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

func True (cmd string) (bool , error){
	return true , nil
}


func main() {
	flag, err := True(os.Args[0])
	if flag {
		os.Exit(0)
	}
	fmt.Println(err) 
}