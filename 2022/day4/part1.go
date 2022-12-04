package main

import (
	"os"
	"strconv"
	"strings"
)

type elf struct {
	area      string
	areas     []string
	areaStart int
	areaEnd   int
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)

	pairs := strings.Split(inputString, "\n")
	result := 0

	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")

		elf1 := elf{area: elfs[0], areas: strings.Split(elfs[0], "-"), areaStart: 0, areaEnd: 0}
		elf2 := elf{area: elfs[1], areas: strings.Split(elfs[1], "-"), areaStart: 0, areaEnd: 0}

		elf1.areaStart, _ = strconv.Atoi(elf1.areas[0])
		elf1.areaEnd, _ = strconv.Atoi(elf1.areas[1])
		elf2.areaStart, _ = strconv.Atoi(elf2.areas[0])
		elf2.areaEnd, _ = strconv.Atoi(elf2.areas[1])

		if elf1.areaStart <= elf2.areaStart && elf1.areaEnd >= elf2.areaEnd {
			result++
		} else if elf2.areaStart <= elf1.areaStart && elf2.areaEnd >= elf1.areaEnd {
			result++
		}
	}
	println(result)
}
