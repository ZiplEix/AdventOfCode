package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func toStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

type monkey struct {
	Name      string
	Items     []string
	operation string
	test      int
	vrai      int
	faux      int
	inspected int
}

func makeMonkey(notes []string) []monkey {
	var monkeys []monkey
	for i := 0; i < 4; i++ {
		monkey := monkey{}
		monkey.Name = notes[i*7][7:8]
		monkey.Items = strings.Split(notes[i*7+1][18:], ", ")
		monkey.operation = notes[i*7+2][13:]
		monkey.test = toInt(notes[i*7+3][21:])
		monkey.vrai = toInt(notes[i*7+4][29:])
		monkey.faux = toInt(notes[i*7+5][30:])
		monkey.inspected = 0
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func makeOperation(operation string, item int) int {
	operation = strings.Replace(operation, "old", toStr(item), -1)
	op := strings.Split(operation, " ")
	op = op[2:]
	switch op[1] {
	case "+":
		// fmt.Println("    Worry level increases by", op[2], "to", toInt(op[0])+toInt(op[2]))
		return toInt(op[0]) + toInt(op[2])
	case "-":
		// fmt.Println("    Worry level decreases by", op[2], "to", toInt(op[0])-toInt(op[2]))
		return toInt(op[0]) - toInt(op[2])
	case "*":
		// fmt.Println("    Worry level multiplied by", op[2], "to", toInt(op[0])*toInt(op[2]))
		return toInt(op[0]) * toInt(op[2])
	case "/":
		// fmt.Println("    Worry level divided by", op[2], "to", toInt(op[0])/toInt(op[2]))
		return toInt(op[0]) / toInt(op[2])
	case "%":
		// fmt.Println("    Worry level % by", op[2], "to", toInt(op[0])/toInt(op[2]))
		return toInt(op[0]) % toInt(op[2])
	}
	return 0
}

func main() {
	input, err := os.ReadFile("./inputTest")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	notes := strings.Split(inputString, "\n")

	monkeys := makeMonkey(notes)

	for k := 0; k < 10000; k++ {
		for i := 0; i < len(monkeys); i++ {
			// fmt.Println("Monkey", monkeys[i].Name, ":")
			for len(monkeys[i].Items) > 0 {
				// fmt.Println("  Monkey inspects an item with a worry level of", monkeys[i].Items[0])
				monkeys[i].inspected++
				currentItem := toInt(monkeys[i].Items[0])
				currentItem = makeOperation(monkeys[i].operation, currentItem)
				// currentItem /= 3
				// fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to ", currentItem)
				if currentItem%monkeys[i].test == 0 {
					// fmt.Println("    Current worry level is divisible by", monkeys[i].test)
					// fmt.Println("    Item with worry level", currentItem, "is thrown to monkey", monkeys[i].vrai)
					monkeys[monkeys[i].vrai].Items = append(monkeys[monkeys[i].vrai].Items, toStr(currentItem))
				} else {
					// fmt.Println("    Current worry level is not divisible by", monkeys[i].test)
					// fmt.Println("    Item with worry level", currentItem, "is thrown to monkey", monkeys[i].faux)
					monkeys[monkeys[i].faux].Items = append(monkeys[monkeys[i].faux].Items, toStr(currentItem))
				}
				monkeys[i].Items = monkeys[i].Items[1:]
			}
		}
		if (k+1)%1000 == 0 || (k+1) == 1 || (k+1) == 20 && k != 0 {
			fmt.Println("== After round", k+1, "==")
			fmt.Println("Monkey 0:", monkeys[0].inspected)
			fmt.Println("Monkey 1:", monkeys[1].inspected)
			fmt.Println("Monkey 2:", monkeys[2].inspected)
			fmt.Println("Monkey 3:", monkeys[3].inspected)
			fmt.Println()
		}
	}

	var inspected []int
	for i := 0; i < len(monkeys); i++ {
		inspected = append(inspected, monkeys[i].inspected)
	}
	sort.Ints(inspected)
	result := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	println(result)
}
