package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("2024/day3/input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt")
		return
	}
	strInput := string(readFile)

	var index int = 0
	solution := 0
	doDontSolution := 0
	shouldConsiderThisMul := true

	for {
		index = strings.Index(strInput, "mul(")
		if index == -1 {
			break
		}
		doIndex := strings.LastIndex(strInput[:index], "do()")
		dontIndex := strings.LastIndex(strInput[:index], "don't()")

		if dontIndex != -1 || doIndex != -1 {
			if doIndex == -1 && dontIndex < index {
				shouldConsiderThisMul = false
			} else if dontIndex == -1 && doIndex < index {
				shouldConsiderThisMul = true
			} else {
				maxIndex := math.Max(float64(dontIndex), float64(doIndex))
				if doIndex < index && index < dontIndex {
					shouldConsiderThisMul = true
				} else if dontIndex < index && index < doIndex {
					shouldConsiderThisMul = false
				} else if int(maxIndex) < index {
					if int(maxIndex) == doIndex {
						shouldConsiderThisMul = true
					} else if int(maxIndex) == dontIndex {
						shouldConsiderThisMul = false
					}
				}
			}
		}

		strInput = strInput[index+4:]
		closeIndex := strings.Index(strInput, ")")
		if closeIndex == -1 {
			break
		}
		interestedStr := strInput[:closeIndex]

		splitString := strings.Split(interestedStr, ",")
		if len(splitString) == 2 {
			if validString(splitString[0]) && validString(splitString[1]) {
				num1, _ := strconv.Atoi(splitString[0])
				num2, _ := strconv.Atoi(splitString[1])
				solution += num1 * num2
				if shouldConsiderThisMul {
					doDontSolution += num1 * num2
				}
			}
		}
	}
	fmt.Println(solution)
	fmt.Println(doDontSolution)
}

func validString(str string) bool {
	for _, char := range str {
		if !(char-'0' >= 0 && char-'0' <= 9) {
			return false
		}
	}
	return true
}
