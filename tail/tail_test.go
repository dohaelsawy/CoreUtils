package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestTail(t *testing.T) {
	result , err := Tail(5,`D:\CodeScaler\input.txt`)
	if err != nil {
		t.Errorf(err.Error())
	}
	input := `Hello Jan
Hello Cone
Hello Hanje
Hello Eren
this line should not be displayed`

	expected := strings.Split(input, "\n")

	if !reflect.DeepEqual(result, expected){
		t.Errorf("we expected %v , we got %v" , expected , result)
	}
	t.Logf("Expected: %v\n", expected)
	t.Logf("Result: %v\n", result)
}