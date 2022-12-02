package main

import (
	"fmt"
	"os"
	"strings"
)

/*
a = pierre
b = feuille
c = ciseaux
x = lose
y = draw
z = win
*/

func definePoints(him, me string) int {
	if him == "A" {
		if me == "X" {
			return 0 + 3
		} else if me == "Y" {
			return 3 + 1
		} else if me == "Z" {
			return 6 + 2
		}
	} else if him == "B" {
		if me == "X" {
			return 0 + 1
		} else if me == "Y" {
			return 3 + 2
		} else if me == "Z" {
			return 6 + 3
		}
	} else if him == "C" {
		if me == "X" {
			return 0 + 2
		} else if me == "Y" {
			return 3 + 3
		} else if me == "Z" {
			return 6 + 1
		}
	}
	return 0
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)

	rounds := strings.Split(inputString, "\n")
	result := 0

	for _, round := range rounds {
		actions := strings.Split(round, " ")
		him := actions[0]
		me := actions[1]

		result += definePoints(him, me)
	}
	fmt.Println(result)
}
