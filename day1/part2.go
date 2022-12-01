// a hello word program

package main

import (
	"os"
	"sort"
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
	elfsArray := make([]int, 0)

	elfs := strings.Split(inputString, "\n\n")
	for _, elf := range elfs {
		calories := strings.Split(elf, "\n")
		elfFood := 0
		for _, calorie := range calories {
			Kca, _ := strconv.Atoi(calorie)
			elfFood += Kca
		}
		elfsArray = append(elfsArray, elfFood)
	}
	sort.Ints(elfsArray)

	for i := len(elfsArray) - 1; i >= len(elfsArray)-3; i-- {
		result += elfsArray[i]
	}

	println(result)
}
