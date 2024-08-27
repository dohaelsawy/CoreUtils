package main

import (
	"fmt"
	"os"
)

func Env() {
	for _,env := range os.Environ() {
		fmt.Println(env)
	} 
}

func main() {
	Env()
}