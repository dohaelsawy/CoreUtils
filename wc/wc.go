package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type ErrTypes string

const (
	ErrTxtNotFound        = ErrTypes("the file you desire is not found :{")
	ErrTxtNotEnoughLines  = ErrTypes("the file doesn't contain 10 lines, you can use -n flag to spacify the number of lines you want")
	ErrCommandNotExist    = ErrTypes("the following command doesn't exist")
	ErrSomethingWentWrong = ErrTypes("sorry something went wrong try again :}")
)

func (e ErrTypes) Error() string {
	return string(e)
}

func WordCount(input string) int{
	return len(strings.Fields(input))
}
func CharCount(input string) int{
	return len(input)
}
func LineCount(input string) int{
	lines := strings.Split( strings.TrimSpace(string(input)), "\n")
	return len(lines)
}

func WC(l,w,c bool, file_name string) (int,int,int, error) {
	data, err := os.ReadFile(file_name)
	if err != nil {
		return 0,0,0,ErrTxtNotFound
	} 
	return LineCount(string(data)), WordCount(string(data)), CharCount(string(data)), nil
}

func main() {
	var w , l , c bool
	flag.BoolVar(&w,"w", false, "last n lines from mentioned file")
	flag.BoolVar(&l,"l", false, "last n lines from mentioned file")
	flag.BoolVar(&c,"c", false, "last n lines from mentioned file")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		filePath := args[0]
		lines ,words ,chars, err := WC(l,w,c, filePath)
		if err == ErrTxtNotFound {
			fmt.Println(err)
		} 
		if !l && !w && !c {
			fmt.Println(lines,words,chars)
		}else {
			if l {
				fmt.Println(lines)
			}
			if w {
				fmt.Println(words)
			}
			if c {
				fmt.Println(chars)
			}
		}

	} else {
		fmt.Println(ErrCommandNotExist)
	}
}
