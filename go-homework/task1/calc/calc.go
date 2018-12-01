package main

import (
	"fmt"
	"strconv"
	"strings"
)

// сюда писать код
// фукция main тоже будет тут

func calc(seq string) (int, error) {
	stack := make([]int, len(seq))
	sp := 0
	array := strings.Split(seq, " ")
	for _, elem := range array {
		switch string(elem) {
		case " ":
		case "\n":
		case "=":
			return stack[sp-1], nil
		case "+":
			stack[sp-2] = stack[sp-2] + stack[sp-1]
			sp = sp - 1
		case "-":
			stack[sp-2] = stack[sp-2] - stack[sp-1]
			sp = sp - 1
		case "*":
			stack[sp-2] = stack[sp-1] * stack[sp-2]
			sp = sp - 1
		case "/":
			stack[sp-2] = stack[sp-2] / stack[sp-1]
			sp = sp - 1

		default:
			val, err := strconv.Atoi(string(elem))
			if err != nil {
				return -1, err
			}
			stack[sp] = val
			sp++
		}
	}
	return stack[sp-1], nil
}

func main() {
	var s string = "0.1 3 / 4 5 * -"
	fmt.Println(calc(s))
}
