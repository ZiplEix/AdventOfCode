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
	result := 0

	elfs := strings.Split(inputString, "\n\n")
	for _, elf := range elfs {
		calories := strings.Split(elf, "\n")
		elfFood := 0
		for _, calorie := range calories {
			Kca, _ := strconv.Atoi(calorie)
			elfFood += Kca
		}
		if elfFood > result {
			result = elfFood
		}
	}

	println(result)
}
