package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func isSortedDesc(slice []int) bool {
	var unsafe int
	for i := 0; i < len(slice)-1; i++ {
		// WHY IS THIS RIGHT???? WHY DID THIS POP OUT A CORRECT RESULT???
		// difference shouldn't be larger than 2
		// moving on
		if slice[i] <= slice[i+1] || slice[i]-slice[i+1] > 3 {
			unsafe++
			if unsafe > 1 {
				return false
			}
		}
	}
	return true
}

func isSortedAsc(slice []int) bool {
	var unsafe int
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] >= slice[i+1] || slice[i+1]-slice[i] > 3 {
			unsafe++
			if unsafe > 1 {
				return false
			}
		}
	}
	return true
}

func reportsArray(fileName string) ([]string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var reports []string

	// just learned abt scanners
	scanner := bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		reports = append(reports, scanner.Text())
	}

	return reports, nil
}

func countSafeReports(reports []string) int {
	var safeReports int = 0
	for i := range reports {
		temp, err := stringsToIntegers(strings.Fields(reports[i]))
		if err != nil {
			log.Fatal(err)
		}
		if isSortedAsc(temp) || isSortedDesc(temp) {
			safeReports++
		}
	}

	return safeReports
}

func main() {
	reports, err := reportsArray("./input.txt")
	if err != nil {
		return
	}
	fmt.Println(countSafeReports(reports))
}
