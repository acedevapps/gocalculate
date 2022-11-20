package main

import (
	"fmt"
	"math"
	"strings"
)

var precedence []string
var operators []string

func operate(op string, a, b float64) float64 {
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
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	operators = []string{"*", "/", "+", "-"}
	precedence = []string{"*/", "+-"}
	fmt.Println(calculate("1.523781*3.23132"))
}

func calculate(equ string) float64 {
	var stack []float64
	var opStack []string
	var cur float64
	eq := strings.Replace(equ, " ", "", -1)
	d := 0
	for _, s := range eq {
		if contains(operators, string(s)) {
			opStack = append(opStack, string(s))
			stack = append(stack, cur)
			cur = 0
			d = 0
		} else if string(s) == "." {
			d++
		} else {
			if d != 0 {
				cur = cur + float64(s-'0')/float64(math.Pow(10, float64(d)))
				d++

			} else {
				cur = cur*10 + float64(s-'0')
			}
		}
	}
	stack = append(stack, cur)
	j := 0
	for _, op := range precedence {
		for index := 0; index < len(opStack); index++ {
			if strings.Contains(op, opStack[index]) {
				i := index - j
				stack[i] = operate(opStack[index], stack[i], stack[i+1])
				if i+1 < len(stack)-1 {
					copy(stack[i+1:], stack[i+2:])
					stack[len(stack)-1] = 0
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
				}
				opStack[index] = "e"
				j++
			}
		}
		for contains(opStack, "e") {
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
		j = 0
	}
	return stack[0]
}
