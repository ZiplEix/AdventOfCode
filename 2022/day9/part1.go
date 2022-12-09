package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func move(xH, yH, xT, yT int) (int, int) {
	if xH == xT && yH == yT {
		return xT, yT
	}

	if (xH-xT == 1 || xH-xT == -1) && yH == yT {
		return xT, yT
	}
	if (yH-yT == 1 || yH-yT == -1) && xH == xT {
		return xT, yT
	}
	if (xH-xT == 1 || xH-xT == -1) && (yH-yT == 1 || yH-yT == -1) {
		return xT, yT
	}

	if xH > xT && yH == yT {
		return xT + 1, yT
	}
	if xH < xT && yH == yT {
		return xT - 1, yT
	}
	if xH == xT && yH > yT {
		return xT, yT + 1
	}
	if xH == xT && yH < yT {
		return xT, yT - 1
	}

	if xH > xT && yH > yT {
		return xT + 1, yT + 1
	}
	if xH > xT && yH < yT {
		return xT + 1, yT - 1
	}
	if xH < xT && yH > yT {
		return xT - 1, yT + 1
	}
	if xH < xT && yH < yT {
		return xT - 1, yT - 1
	}

	return xT, yT
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	steps := strings.Split(inputString, "\n")

	carte := map[string]int{"0;0": 1}
	xH, yH, xT, yT := 0, 0, 0, 0

	for _, step := range steps {
		direction := string(step[0])
		distance, _ := strconv.Atoi(step[2:])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				yH++
				xT, yT = move(xH, yH, xT, yT)
				carte[strconv.Itoa(xT)+";"+strconv.Itoa(yT)] = 1
			case "D":
				yH--
				xT, yT = move(xH, yH, xT, yT)
				carte[strconv.Itoa(xT)+";"+strconv.Itoa(yT)] = 1
			case "R":
				xH++
				xT, yT = move(xH, yH, xT, yT)
				carte[strconv.Itoa(xT)+";"+strconv.Itoa(yT)] = 1
			case "L":
				xH--
				xT, yT = move(xH, yH, xT, yT)
				carte[strconv.Itoa(xT)+";"+strconv.Itoa(yT)] = 1
			}
		}
	}
	fmt.Println(len(carte))
}
