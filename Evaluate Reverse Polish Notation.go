package Leetcode

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int

	for _, rec := range tokens {
		switch rec {
		case "+":
			if len(stack) < 2 {
				panic("length of stack error")
			}
			tmp := stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack,tmp)
		case "-":
			if len(stack) < 2 {
				panic("length of stack error")
			}
			tmp := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack,tmp)
		case "*":
			if len(stack) < 2 {
				panic("length of stack error")
			}
			tmp := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack,tmp)
		case "/":
			if len(stack) < 2 {
				panic("length of stack error")
			}
			tmp := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack,tmp)
		default:
			data, _ := strconv.Atoi(rec)
			stack = append(stack, data)
		}
	}

	return stack[0]
}
