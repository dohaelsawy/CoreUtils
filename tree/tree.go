package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)


type ErrTypes string

const (
	ErrFileNotFound = ErrTypes("the file you desire is not found :{")
	ErrCommandNotExist = ErrTypes("the following command doesn't exist")
	ErrSomethingWentWrong = ErrTypes("sorry something went wrong try again :}")
)


func (e ErrTypes) Error() string {
	return string(e)
}

func TreeRecursive(currentLevel int, path , spaces string){
	if currentLevel < 1 {
		return
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Print(dirs)
		fmt.Println(ErrFileNotFound)
		return
	}
	if len(dirs) == 0 {
		return 
	}
	currentLevel--
	for _, dir := range dirs {
		if dir.IsDir(){
			fmt.Printf("%s%s\n", spaces ,dir.Name())
			TreeRecursive(currentLevel, filepath.Join(path, dir.Name()), spaces+"    ")
		}else {
			fmt.Printf("%s%s\n", spaces,dir.Name())
		}
	}
}

func main() {
	var l int
	flag.IntVar(&l,"L", 1, "last n lines from mentioned file")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		TreeRecursive(l,args[len(args)-1],"")
	}
}
