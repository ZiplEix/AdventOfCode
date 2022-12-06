package main

import (
	"fmt"
	"os"
	"strings"
)

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func isMarker(str string) bool {
	leters := strings.Split(str, "")
	leters = removeDuplicateStr(leters)
	fmt.Println(leters)
	if len(leters) != 4 {
		return false
	}
	return true
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)

	for i := 0; i < len(inputString)-4; i++ {
		if isMarker(inputString[i : i+4]) {
			fmt.Println(i + 4)
			break
		}
	}
}
