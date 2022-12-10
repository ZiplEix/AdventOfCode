package main

import (
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
	instructions := strings.Split(inputString, "\n")

	cycle := 1
	X := 1
	instructionIndex := 0
	addCycle := 0

	for cycle <= 240 {
		if instructions[instructionIndex] == "noop" {
			instructionIndex++
		} else if instructions[instructionIndex][:4] == "addx" {
			if addCycle == 0 {
				addCycle = 1
			} else if addCycle == 1 {
				nb, _ := strconv.Atoi(instructions[instructionIndex][5:])
				X += nb
				addCycle = 0
				instructionIndex++
			}
		}

		cycle++

		if (cycle-1)%40 == X || (cycle-1)%40 == X-1 || (cycle-1)%40 == X+1 {
			print("#")
		} else {
			print(".")
		}

		if (cycle-1)%40 == 0 {
			print("\n")
		}
	}
}
