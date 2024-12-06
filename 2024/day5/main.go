package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var orderMap = make(map[int][]int)

func main() {

	file, err := os.ReadFile("2024/day5/input.txt")
	if err != nil {
		return
	}
	lines := strings.Split(string(file), "\n")

	for idx, line := range lines {
		if line == "" {
			fmt.Println(CountValidUpdates(lines[idx+1:]))
			return
		}
		strSplit := strings.Split(string(line), "|")
		num1, _ := strconv.Atoi(strSplit[0])
		num2, _ := strconv.Atoi(strSplit[1])
		if val, ok := orderMap[num1]; ok {
			val = append(val, num2)
			orderMap[num1] = val
		} else {
			orderMap[num1] = []int{num2}
		}
	}
}

func CountValidUpdates(updates []string) int {
	solution := 0

	for _, update := range updates {
		updateSlice := getSliceOfIntegers(update)
		solution += findIftheUpdateIsValid(updateSlice)
	}
	return solution
}

// findIftheUpdateIsValid returns the middle number if the update is valid or returns 0
func findIftheUpdateIsValid(slice []int) int {
	for idx1, num1 := range slice {
		for idx2 := idx1 + 1; idx2 < len(slice); idx2++ {
			if slices.Index(orderMap[num1], slice[idx2]) == -1 {
				return 0
			}
		}
	}
	return slice[len(slice)/2]
}

func getSliceOfIntegers(update string) []int {
	var updateSlice []int
	strSplit := strings.Split(update, ",")
	for _, str := range strSplit {
		num1, _ := strconv.Atoi(str)
		updateSlice = append(updateSlice, num1)
	}
	return updateSlice
}
