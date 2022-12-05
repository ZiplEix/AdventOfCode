package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)

	stackData := strings.Split(inputString, "\n")[:8]
	
	actions := strings.Split(inputString, "\n")[10:]

	stack := make([][]string, 9, 9)
	for _, line := range stackData {
		for i, box := range line {
			if i%4 == 1 && box != ' ' {
				stack[i/4] = append([]string{string(box)}, stack[i/4]...)
			}
		}
	}

	for _, action := range actions {
		action = strings.ReplaceAll(action, "move ", "")
		action = strings.ReplaceAll(action, "from ", "")
		action = strings.ReplaceAll(action, "to ", "")

		action := strings.Split(action, " ")
		nb, _ := strconv.Atoi(action[0])
		from, _ := strconv.Atoi(action[1])
		to, _ := strconv.Atoi(action[2])

		fmt.Println(nb, from, to)

		for i := 0; i < nb; i++ {
			stack[to-1] = append(stack[to-1], stack[from-1][len(stack[from-1])-1])
			stack[from-1] = stack[from-1][:len(stack[from-1])-1]
		}
	}

	for _, column := range stack {
		fmt.Print(column[len(column)-1])
	}
	fmt.Println()
}
