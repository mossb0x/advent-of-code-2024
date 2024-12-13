package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInput(filename string) ([][]int, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var reports [][]int

	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	for scanner.Scan() {
		row := []int{}
		for _, value := range strings.Fields(scanner.Text()) {
			temp, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			row = append(row, temp)
		}
		reports = append(reports, row)
	}

	return reports, nil
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func isSafe(report []int) bool {
	if report[0] < report[1] {
		for i := 0; i < len(report)-1; i++ {
			if report[i+1]-report[i] > 3 || report[i+1]-report[i] < 1 || report[i+1]-report[i] < 0 {
				return false
			}
		}
	} else {
		for i := 0; i < len(report)-1; i++ {
			if report[i]-report[i+1] > 3 || report[i]-report[i+1] < 1 || report[i]-report[i+1] < 0 {
				return false
			}
		}
	}
	return true
}

// The issue here is its not checking if the array is valid without the first number
// FIXME: check if the first number invalidates a report

func isSafeDampener(report []int) bool {
	if report[0] == report[1] {
		return isSafe(slices.Delete(report, 1, 2))
	}

	if report[0] < report[1] {
		for i := 0; i < len(report)-1; i++ {
			if report[i+1]-report[i] > 3 || report[i+1]-report[i] < 1 || report[i+1]-report[i] < 0 {
				return isSafe(slices.Delete(report, i+1, i+2))
			}
		}
	} else {
		for i := 0; i < len(report)-1; i++ {
			if report[i]-report[i+1] > 3 || report[i]-report[i+1] < 1 || report[i]-report[i+1] < 0 {
				return isSafe(slices.Delete(report, i+1, i+2))
			}
		}
	}
	return true
}

func countSafeReports(reports [][]int, dampener bool) int {
	safeReports := 0
	if dampener {
		for _, row := range reports {
			if isSafeDampener(row) {
				safeReports++
			}
		}
	} else {
		for _, row := range reports {
			if isSafe(row) {
				safeReports++
			}
		}
	}
	return safeReports
}

func main() {
	reports, err := getInput("./day2/input.txt")
	if err != nil {
		return
	}

	// reports_test := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// }
	//
	// fmt.Println(countSafeReports(reports, false))
	// fmt.Println(reports)
	fmt.Println("Without dampener")
	fmt.Println(countSafeReports(reports, false))
	fmt.Println("With dampener")
	fmt.Println(countSafeReports(reports, true))
}
