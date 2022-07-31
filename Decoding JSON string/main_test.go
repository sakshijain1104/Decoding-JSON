// This is the testing package

package main

import (
	"fmt"
	"testing"
)

//TestDecodingJson tests whether the Reading variable updated in DecodingJson function (main package) gives the correct output or not
func TestDecodingJson(t *testing.T) {
	expectedOutput := Employee{"Mark Taylor", "Software Developer", "123-466-799", "markt@gmail.com"}
	DecodingJson()

	fmt.Printf("Type of Reading %T\n", Reading)
	fmt.Println(Reading)
	fmt.Printf("Type of Reading %T\n", expectedOutput)
	fmt.Println(expectedOutput)

	// If the expected output values do not match those of the Reading variable, we get an error
	if Reading != expectedOutput {
		t.Error("Not the expected output")
		
	}
}
