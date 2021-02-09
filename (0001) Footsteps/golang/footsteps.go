package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

const DIR_PATH = "footsteps/golang/"
const RIGHT_FEET_OFFSET = 2
const FEET_SAMPLE_FILEPATH = DIR_PATH + "sample.txt"

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
	file, err := os.Create(path)
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

func mirrorRow(s string) string {
	r := []rune(s)
	for i := 0; i <= len(r) / 2; i++ {
		r[len(r) - 1 - i], r[i] = r[i], r[len(r) - 1 - i]
	}
	return string(r)
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

		for j := 0; j < len(sample[i]); j++ {
			symbol := string(sample[i][j])
			mirroredSymbol := MIRRORING_MAP[symbol]
			if mirroredSymbol != "" {
				symbol = mirroredSymbol
			}
			builder.WriteString(symbol)
		}
		mirroredSample[i] = mirrorRow(builder.String())
	}

	return mirroredSample
}

func getRow(step int, stepCnt int, index int, leftFeet []string, rightFeet []string) string {
	var builder strings.Builder
	builder.Grow(len(leftFeet) + len(rightFeet))

	var leftPart string
	if step != 0 && step != stepCnt {
		leftIndex := (index + RIGHT_FEET_OFFSET) % len(leftFeet)
		leftPart = leftFeet[leftIndex]
	} else {
		leftPart = strings.Repeat(" ", len(leftFeet) + 1)
	}

	builder.WriteString(leftPart)
	builder.WriteString(rightFeet[index])

	return builder.String()
}

func printFootsteps(filepath string, personCnt int, stepCnt int) {
	rowIndex := 0

	leftFeet, _ := readLines(FEET_SAMPLE_FILEPATH)
	rightFeet := mirrorFeetSample(leftFeet)

	var lines []string

	for step := 0; step <= stepCnt; {
		var builder strings.Builder
		builder.Grow(personCnt * (len(leftFeet) + len(rightFeet)))

		for person := 0; person < personCnt; person++ {
			row := getRow(step, stepCnt, rowIndex, leftFeet, rightFeet)
			builder.WriteString(row)
		}
		lines = append(lines, builder.String())

		iterateSample(&rowIndex, len(leftFeet))
		if rowIndex == RIGHT_FEET_OFFSET || rowIndex == 0 {
			step += 1
		}
	}


	err := writeLines(lines, filepath)
	if err != nil {
		panic(err)
	}
}

func main() {
	path := flag.String("path", DIR_PATH + "feet.txt", "Путь до результирующего файла")
	personCnt := flag.Int("p", 3, "Количество человек")
	stepCnt := flag.Int("s", 5, "Количество шагов")

	flag.Parse()
	printFootsteps(*path, *personCnt, *stepCnt)
}
