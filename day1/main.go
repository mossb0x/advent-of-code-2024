package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// takes the name of a file to read the inputs of the puzzle from inside it
func separateLists(fileName string) ([]int, []int) {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Couldn't read input.txt")
	}

	rawLists := string(f)
	// turns it into an array with the format XXXXX XXXXX for each element value as string
	lessRawLists := strings.Split(rawLists, "\n")
	var separatedList1, separatedList2 []int
	// iterates over each element
	for i := range lessRawLists {
		// checks for empty string and continues if its there so the program doesnt panic
		if lessRawLists[i] == "" {
			continue
		}

		// converts elements to integers after splitting them with Fields so we can append
		// them to the final lists and use them in listDistance
		temp1, err := strconv.Atoi(strings.Fields(lessRawLists[i])[0])
		if err != nil {
			log.Fatal("ERROR")
		}

		temp2, err := strconv.Atoi(strings.Fields(lessRawLists[i])[1])
		if err != nil {
			log.Fatal("ERROR")
		}

		separatedList1 = append(separatedList1, temp1)
		separatedList2 = append(separatedList2, temp2)
	}

	return separatedList1, separatedList2
}

func listDistance(list1 []int, list2 []int) int {
	// sorts the lists
	sort.Ints(list1)
	sort.Ints(list2)
	// we store the distance here
	var distance int = 0

	var distanceList []int

	for i, value1 := range list1 {
		value2 := list2[i]
		// i hate how this looks :/
		distanceList = append(distanceList, int(math.Abs(float64(value1-value2))))
	}

	for _, value := range distanceList {
		distance += value
	}

	return distance
}

func main() {
	list1, list2 := separateLists("input.txt")
	fmt.Println(listDistance(list1, list2))
}
