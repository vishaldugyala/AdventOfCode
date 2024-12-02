package utils

import (
	"bufio"
	"os"
)

func GetInputFromFile(filename string) ([][]byte, error) {
	open, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	reader := bufio.NewReader(open)
	var linesRead [][]byte
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		linesRead = append(linesRead, line)
	}
	return linesRead, nil

}
