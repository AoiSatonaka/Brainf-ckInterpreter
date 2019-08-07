package bfi

import (
	"errors"
	"fmt"
)

func Interpret(sourceCode []int32, input []int32) error {
	var cache [65535]int32
	p := 0
	ls := 0

	for si := 0; si < len(sourceCode); si++ {
		order := string(sourceCode[si])

		switch order {
		case "+":
			cache[p]++
		case "-":
			cache[p]--
		case ">":
			p++
		case "<":
			p--
		case ".":
			fmt.Print(string(cache[p]))
		case ",":
			cache[p] = input[0]
			input = input[1:]
		case "[":
			ls = si
		case "]":
			if cache[p] > 0 {
				si = ls
			}
		default:
			return errors.New(order + " is impossible to analyze")
		}
	}
	return nil
}
