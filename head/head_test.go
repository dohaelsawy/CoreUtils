package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestHead(t *testing.T) {
	result , err := Head(5,`D:\CodeScaler\input.txt`)
	if err != nil {
		t.Errorf(err.Error())
	}
	input := `Hello CodeScalers
Hello world
Hello Sasha
Hello Mickasa
Hello Levi`

	expected := strings.Split(input, "\n")

	if !reflect.DeepEqual(result, expected){
		t.Errorf("we expected %v , we got %v" , expected , result)
	}
	t.Logf("Expected: %v\n", expected)
	t.Logf("Result: %v\n", result)
}