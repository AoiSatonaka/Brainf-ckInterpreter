package Brainf_ckInterpreter

import (
	"fmt"
)


func Interpret(sourceCode []int32, input []int32) error {
	var cache [65535]int32
	p := 0
	ls := 0

	for si:=0; si<len(sourceCode); si++ {
		order := string(sourceCode[si])

		switch order {
		case "+": cache[p]++
		case "-": cache[p]--
		case ">": p++
		case "<": p--
		case ".": fmt.Println(string(cache[p]))
		case ",":
			cache[p] = input[0]
			input = input[1:]
		case "[": ls = si
		case "]": if cache[p] > 0 { si = ls }
		}
		//fmt.Println("pointer:",p)
		//fmt.Println("cache[p]",cache[p])
	}
	return nil
}

