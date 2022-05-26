package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var counter int
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "+" || string(input[i]) == "-" {
			counter++
		}
	}

	if input == " " {
		output = " "
		err = errorEmptyInput
	} else if counter != 2 {
		output = " "
		err = errorNotTwoOperands
	} else if input != " " && counter == 2 {
		var cnt int
		var num, sum int64
		var tmpstr string
		input = strings.ReplaceAll(input, " ", "")

		for pos := 1; pos < len(input)+1; pos++ {
			for i := pos; i < len(input) && input[i] >= '0' && input[i] <= '9'; i++ {
				cnt++
			}

			tmpstr = input[pos-1 : pos+cnt]
			num, _ = strconv.ParseInt(tmpstr, 10, 64)
			sum += num
			pos += cnt
			cnt = 0
		}
		temp_sum := int(sum)

		output = strconv.Itoa(temp_sum)

	}
	//fmt.Println(counter)
	return output, err

}

func main() {
	fmt.Println(StringSum("25+6+7+8"))
}
