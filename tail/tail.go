package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

func Tail(n int,file_name string) ([]string,error) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		return []string{}, ErrTxtNotFound
	}else {
		lines := strings.Split(string(data), "\n")
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}
		if len(lines) < n {
			return lines , ErrTxtNotEnoughLines
		}
		return lines[len(lines)-n:], nil
	}
}
func main() {
	n := flag.Int("n", 10, "last n lines from mentioned file")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		filePath := args[0]
		lines, err := Tail(*n, filePath)
		if err == ErrTxtNotFound {
			fmt.Println(err)
		} else {
			for _, line := range lines {
				fmt.Println(line)
			}
			if err == ErrTxtNotEnoughLines {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(ErrCommandNotExist)
	}
}

