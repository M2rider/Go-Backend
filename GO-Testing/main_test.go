package main

import(
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T){
	if Calculate(2)!=4 {
		t.Error("Expected 2+2 equal to 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct{
		input int
		expected int
	} {
		{2,4},
		{-1,1},
		{0,2},
		{9999,10001},
	}

	for _,test := range tests {
		output := Calculate(test.input)
		if output != test.expected {
			t.Errorf("Test Failed: {%v} inputted, {%v} expected, received: {%v}",test.input,test.expected,output)
		}else {
			fmt.Printf("Test Passed: {%v} inputted, {%v} expected, received: {%v}\n",test.input,test.expected,output)
		}
	}
}