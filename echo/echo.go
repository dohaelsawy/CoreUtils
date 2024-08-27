package main

import (
	"flag"
	"fmt"
)
type ErrTypes string

const (
	ErrTxtNotFound = ErrTypes("the file you desire is not found :{")
	ErrTxtNotEnoughLines = ErrTypes("the file doesn't contain 10 lines, you can use -n flag to spacify the number of lines you want")
	ErrCommandNotExist = ErrTypes("the following command doesn't exist")
	ErrSomethingWentWrong = ErrTypes("sorry something went wrong try again :}")
)


func (e ErrTypes) Error() string {
	return string(e)
}


func Echo(n bool, output string) (error){
	if n {
		fmt.Print(output)
		return nil
	}
	fmt.Println(output)
	return nil
}

func main() {
	n := flag.Bool("n", false, "omit the trailing newline")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		err := Echo(*n ,args[len(args)-1]) 
		if err != nil{
			fmt.Println(err)
		}
	}else {
		fmt.Println(ErrCommandNotExist)
	}
}