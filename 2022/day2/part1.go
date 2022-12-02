package main

import (
	"fmt"
	"os"
	"strings"
)

/*
a = x = pierre
b = y = feuille
c = z = ciseaux
*/

func definePoints(him, me string) int {
	if him == "A" {
		if me == "X" {
			return 1 + 3
		} else if me == "Y" {
			return 2 + 6
		} else if me == "Z" {
			return 3 + 0
		}
	} else if him == "B" {
		if me == "X" {
			return 1 + 0
		} else if me == "Y" {
			return 2 + 3
		} else if me == "Z" {
			return 3 + 6
		}
	} else if him == "C" {
		if me == "X" {
			return 1 + 6
		} else if me == "Y" {
			return 2 + 0
		} else if me == "Z" {
			return 3 + 3
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
