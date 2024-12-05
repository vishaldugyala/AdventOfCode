package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.ReadFile("2024/day2/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	var reports [][]int
	strSlice := strings.Split(string(readFile), "\n")
	for _, str := range strSlice {
		strSlice2 := strings.Split(str, " ")
		var tempSlice []int
		for _, str2 := range strSlice2 {
			tempInt, _ := strconv.Atoi(str2)
			tempSlice = append(tempSlice, tempInt)
		}
		reports = append(reports, tempSlice)
	}

	solution := 0

	for _, report := range reports {
		solution += isTheReportSafe(report)
	}
	fmt.Println(solution)

	solution = 0
	for _, report := range reports {
		solution += isTheReportSafeAfterRemovingASingleLevel(report)
	}
	fmt.Println(solution)
}

func isTheReportSafeAfterRemovingASingleLevel(report []int) int {

	if len(report) < 2 {
		return 1
	}

	//is the report safe after removing first element
	if isTheReportSafe(report) == 1 {
		return 1
	}

	for i := 0; i < len(report); i++ {
		tempReport := slices.Clone(report)
		tempReport = slices.Delete(tempReport, i, i+1)
		if isTheReportSafe(tempReport) == 1 {
			return 1
		}
	}

	return 0
}

func findNumberOfSafeReports(input [][]byte) int {
	reports := getReports(input)
	solution := 0

	for _, report := range reports {
		solution += isTheReportSafe(report)
	}
	return solution
}

func isTheReportSafe(report []int) int {
	if len(report) < 2 {
		return 1
	}
	sign := -1

	for i := 0; i < len(report)-1; i++ {
		if i == 0 {
			if report[i+1] > report[i] {
				sign = 1
			}
		}
		if sign == 1 && ((report[i+1]-report[i] > 3) || (report[i+1] <= report[i])) {
			return 0
		} else if sign == -1 && ((report[i]-report[i+1] > 3) || (report[i+1] >= report[i])) {
			return 0
		}
	}
	return 1
}

func getReports(input [][]byte) [][]int {
	var reports [][]int
	for _, input := range input {
		var tempReport []int
		strSlice := strings.Split(string(input), " ")
		for _, str := range strSlice {
			level, _ := strconv.Atoi(str)
			tempReport = append(tempReport, level)
		}
		reports = append(reports, tempReport)
	}
	return reports
}
