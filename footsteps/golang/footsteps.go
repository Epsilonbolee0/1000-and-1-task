package golang

import (
	"bufio"
	_ "fmt"
	"os"
	"strings"
)

const RIGHT_FEET_OFFSET = 2
const FEET_SAMPLE_FILEPATH = "sample.txt"

var MIRRORING_MAP = map[string]string{
	"(":  ")",
	")":  "(",
	"\\": "/",
	"/":  "\\",
}

func handleClose(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer handleClose(file)
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	return file.Close()
}

func iterateSample(index *int, feetLength int) {
	*index = (*index + 1) % feetLength
}

func mirrorFeetSample(sample []string) []string {
	longestRow := 0
	for _, row := range sample {
		if len(row) > longestRow {
			longestRow = len(row)
		}
	}

	mirroredSample := make([]string, len(sample))

	for i := 0; i < len(sample); i++ {
		var builder strings.Builder
		builder.Grow(longestRow)

		for j := 0; j < longestRow; j++ {
			symbol := sample[i][j:j]
			mirroredSymbol := MIRRORING_MAP[symbol]
			if mirroredSymbol != "" {
				symbol = mirroredSymbol
			}
			builder.WriteString(symbol)
		}
		mirroredSample[i] = builder.String()
	}

	return mirroredSample
}

func getRow(index int, leftFeet []string, rightFeet []string) string {
	var builder strings.Builder
	builder.Grow(len(leftFeet) + len(rightFeet))

	builder.WriteString(leftFeet[index])

	rightIndex := (index + RIGHT_FEET_OFFSET) % len(leftFeet)
	builder.WriteString(rightFeet[rightIndex])
	return builder.String()
}

func printFootsteps(filepath string, personCnt int, stepCnt int) {
	rowIndex := 0

	leftFeet, _ := readLines(FEET_SAMPLE_FILEPATH)
	rightFeet := mirrorFeetSample(leftFeet)

	var lines []string
	for step := 0; step < stepCnt; {

		for person := 0; person < personCnt; {
			row := getRow(rowIndex, leftFeet, rightFeet)
			lines = append(lines, row)
		}
		iterateSample(&rowIndex, len(leftFeet))
	}

	err := writeLines(lines, filepath)
	if err != nil {
		panic(err)
	}
}

func main() {
	printFootsteps("foot.txt", 3, 5)
}
