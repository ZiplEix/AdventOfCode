package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isVisible(rows []string, i, j int) int {
	weight, _ := strconv.Atoi(string(rows[i][j]))

	ib := i - 1
	viewHaut := 0
	for ib >= 0 {
		if nb, _ := strconv.Atoi(string(rows[ib][j])); nb >= weight {
			viewHaut++
			break
		}
		viewHaut++
		ib--
	}
	ib = i + 1
	viewBas := 0
	for ib < len(rows) {
		if nb, _ := strconv.Atoi(string(rows[ib][j])); nb >= weight {
			viewBas++
			break
		}
		viewBas++
		ib++
	}

	jb := j - 1
	viewGauche := 0
	for jb >= 0 {
		if nb, _ := strconv.Atoi(string(rows[i][jb])); nb >= weight {
			viewGauche++
			break
		}
		viewGauche++
		jb--
	}
	jb = j + 1
	viewDroite := 0
	for jb < len(rows[i]) {
		if nb, _ := strconv.Atoi(string(rows[i][jb])); nb >= weight {
			viewDroite++
			break
		}
		viewDroite++
		jb++
	}

	return viewHaut * viewBas * viewGauche * viewDroite
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)

	rows := strings.Split(inputString, "\n")
	result := 0

	i := 1
	for i < len(rows)-1 {
		j := 1
		for j < len(rows[i])-1 {
			if view := isVisible(rows, i, j); view > result {
				result = view
			}
			j++
		}
		i++
	}

	fmt.Println(result)
}
