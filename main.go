package main

// import the required packages
import (
	"fmt"
	"math"
	"strings"
)

// define the precedence of operators and the usable operators
var precedence []string
var operators []string

func operate(op string, a, b float64) float64 {
	// function to do the operations between numbers based on the operator
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func contains(s []string, str string) bool {
	// function to check if a string is in a slice
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	// define the precedence of operators and the usable operators
	operators = []string{"*", "/", "+", "-"}
	precedence = []string{"*/", "+-"}
	// calculate the equation "756/800*9" and print the result
	fmt.Println(calculate("756/800*9"))
}

func calculate(equ string) float64 {
	// define the numerator and operator stacks
	var stack []float64
	var opStack []string
	// define variable which stores the current number
	var cur float64
	// remove whitespaces from the equation
	eq := strings.Replace(equ, " ", "", -1)
	// define variable determing decimal place
	d := 0
	// loop through the equation
	for _, s := range eq {
		if contains(operators, string(s)) {
			// if the current character is an operator
			// append the current number to the numerator stack and the operator to the operator stack
			opStack = append(opStack, string(s))
			stack = append(stack, cur)
			// reset the current number and decimal place
			cur = 0
			d = 0
		} else if string(s) == "." {
			// if the current character is a decimal point register it
			d++
		} else {
			// if the current character is a number and no floating point has been registered
			// add it to the current number*10
			if d != 0 {
				// if the current character is a number and a floating point has been registered
				// add it to the current number*10^(-decimal place)
				cur = cur + float64(s-'0')/float64(math.Pow(10, float64(d)))
				d++
			} else {
				cur = cur*10 + float64(s-'0')
			}
		}
	}
	// append the last number to the numerator stack
	stack = append(stack, cur)
	// define variable storing the amount of performed operations in a precedence level
	j := 0
	// loop through the precedence levels
	for _, op := range precedence {
		// loop through all operators to find the ones in the current precedence level
		for index := 0; index < len(opStack); index++ {
			if strings.Contains(op, opStack[index]) {
				// if the operator is in the current precedence level
				// find the location of the numbers in the numerator stack using index-j
				i := index - j
				// perform the operation between the numbers and replace the first number with the result
				stack[i] = operate(opStack[index], stack[i], stack[i+1])
				// remove the second number from the numerator stack
				if i+1 < len(stack)-1 {
					copy(stack[i+1:], stack[i+2:])
					stack[len(stack)-1] = 0
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
				}
				// replace the operator with a nil value
				opStack[index] = "e"
				// register the performed operation
				j++
			}
		}
		// clear the performed operations from the current precedence level from the operator stack
		// do this by looping through till no more nil characters are found
		for contains(opStack, "e") {
			// loop through stack and remove nil values
			for i, v := range opStack {
				if v == "e" {

					if i != len(opStack)-1 {
						copy(opStack[i:], opStack[i+1:])
						opStack[len(opStack)-1] = ""
						opStack = opStack[:len(opStack)-1]
					} else {
						opStack = opStack[:len(opStack)-1]
					}

				}
			}
		}
		// reset the amount of performed operations in precedence level
		j = 0
	}
	// finally there will only be one number left in the numerator stack
	// return this number as the result
	return stack[0]
}
