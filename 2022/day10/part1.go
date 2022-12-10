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
	instructions := strings.Split(inputString, "\n")

	cycle := 1
	X := 1
	instructionIndex := 0
	addCycle := 0
	result := 0

	for cycle < 220 {
		if instructions[instructionIndex] == "noop" {
			instructionIndex++
			cycle++
		} else if instructions[instructionIndex][:4] == "addx" {
			if addCycle == 0 {
				addCycle = 1
				cycle++
			} else if addCycle == 1 {
				cycle++
				nb, _ := strconv.Atoi(instructions[instructionIndex][5:])
				X += nb
				addCycle = 0
				instructionIndex++
			}
		}
		if cycle == 20 {
			result += X * 20
			fmt.Println("X :", X, "lenght :", X*20, " 20  result :", result)
		} else if cycle == 60 {
			result += X * 60
			fmt.Println("X :", X, "lenght :", X*60, "60  result :", result)
		} else if cycle == 100 {
			result += X * 100
			fmt.Println("X :", X, "lenght :", X*100, "100 result :", result)
		} else if cycle == 140 {
			result += X * 140
			fmt.Println("X :", X, "lenght :", X*140, "140 result :", result)
		} else if cycle == 180 {
			result += X * 180
			fmt.Println("X :", X, "lenght :", X*180, "180 result :", result)
		} else if cycle == 220 {
			result += X * 220
			fmt.Println("X :", X, "lenght :", X*220, "220 result :", result)
			break
		}
	}
}
