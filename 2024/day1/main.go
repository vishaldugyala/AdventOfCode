package main

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {

	lines, err := utils.GetInputFromFile("2024/day1/input.txt")
	if err != nil {
		fmt.Println(err)
		panic("unable to open input")
	}
	var list1 []int
	var list2 []int
	for _, line := range lines {
		strSlices := strings.Split(string(line), " ")
		int1, _ := strconv.Atoi(strSlices[0])
		int2, _ := strconv.Atoi(strSlices[len(strSlices)-1])
		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}
	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Printf("total distance is :%v", getTotalDistance(list1, list2))
	fmt.Printf("the similarity score is :%v", getSimilarityScore(list1, list2))

}

func getTotalDistance(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	solution := 0

	for idx, _ := range list1 {
		solution += int(math.Abs(float64(list1[idx] - list2[idx])))
	}
	return solution
}

func getSimilarityScore(list1, list2 []int) int {

	similarity := 0

	for _, num1 := range list1 {
		counter := 0
		for _, num2 := range list2 {
			if num1 == num2 {
				counter += 1
			}
		}
		similarity += num1 * counter
	}
	return similarity
}
