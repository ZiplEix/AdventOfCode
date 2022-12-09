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
	xH, yH, x1, y1, x2, y2, x3, y3, x4, y4, x5, y5, x6, y6, x7, y7, x8, y8, x9, y9 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	for _, step := range steps {
		direction := string(step[0])
		distance, _ := strconv.Atoi(step[2:])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				yH++
			case "D":
				yH--
			case "R":
				xH++
			case "L":
				xH--
			}
			x1, y1 = move(xH, yH, x1, y1)
			x2, y2 = move(x1, y1, x2, y2)
			x3, y3 = move(x2, y2, x3, y3)
			x4, y4 = move(x3, y3, x4, y4)
			x5, y5 = move(x4, y4, x5, y5)
			x6, y6 = move(x5, y5, x6, y6)
			x7, y7 = move(x6, y6, x7, y7)
			x8, y8 = move(x7, y7, x8, y8)
			x9, y9 = move(x8, y8, x9, y9)
			carte[strconv.Itoa(x9)+";"+strconv.Itoa(y9)] = 1
		}
	}
	fmt.Println(len(carte))
}
