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

func Cat(n bool,file_name string) ([]string,error) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		return []string{}, ErrTxtNotFound
	}
	lines := strings.Split( strings.TrimSpace(string(data)), "\n")
	if n {
		for i := range lines {
			lines[i] = fmt.Sprintf("%d %s", i+1, strings.TrimSpace(lines[i]))
		}
	}else {
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}
	}
	
	return lines, nil
}
func main() {
	n := flag.Bool("n", false, "numbered the lines of mentioned file")
	flag.Parse()
	args := flag.Args()
	
	if len(args) > 0 {
		filePath := args[len(args)-1]
		lines, err := Cat(*n, filePath)
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
		fmt.Println(ErrSomethingWentWrong)
	}
}

