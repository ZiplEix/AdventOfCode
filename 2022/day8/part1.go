package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isVisible(rows []string, i, j int) bool {
	weight, _ := strconv.Atoi(string(rows[i][j]))
	if i == 0 || i == len(rows)-1 || j == 0 || j == len(rows[i])-1 {
		return true
	}

	ib := i - 1
	haut := true
	for ib >= 0 {
		if nb, _ := strconv.Atoi(string(rows[ib][j])); nb >= weight {
			haut = false
			break
		}
		ib--
	}
	ib = i + 1
	bas := true
	for ib < len(rows) {
		if nb, _ := strconv.Atoi(string(rows[ib][j])); nb >= weight {
			bas = false
			break
		}
		ib++
	}

	jb := j - 1
	gauche := true
	for jb >= 0 {
		if nb, _ := strconv.Atoi(string(rows[i][jb])); nb >= weight {
			gauche = false
			break
		}
		jb--
	}
	jb = j + 1
	droite := true
	for jb < len(rows[i]) {
		if nb, _ := strconv.Atoi(string(rows[i][jb])); nb >= weight {
			droite = false
			break
		}
		jb++
	}

	if haut == true || bas == true || gauche == true || droite == true {
		return true
	} else {
		return false
	}
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)

	rows := strings.Split(inputString, "\n")
	result := len(rows)*2 + (len(rows[0])-2)*2

	i := 1
	for i < len(rows)-1 {
		j := 1
		for j < len(rows[i])-1 {
			if isVisible(rows, i, j) {
				result++
			}
			j++
		}
		i++
	}

	fmt.Println(result)
}
