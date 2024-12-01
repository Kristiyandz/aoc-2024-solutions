package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

func main() {

	// PART ONE
	file, err := os.Open("input.txt")
	check(err)

	defer file.Close()

	sc := bufio.NewScanner(file)

	listOne := make([]int, 0)
	listTwo := make([]int, 0)

	for sc.Scan() {
		// read each line and split it
		line := sc.Text()

		// skip empty lines
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "   ")

		// convert each part to to int
		listOneInt, listOneErr := strconv.Atoi(parts[0])
		check(listOneErr)
		listTwoInt, listTwoErr := strconv.Atoi(parts[1])
		check(listTwoErr)

		// create two lists
		listOne = append(listOne, listOneInt)
		listTwo = append(listTwo, listTwoInt)

	}

	if err := sc.Err(); err != nil {
		log.Fatal("error scanning the file")
	}

	// sort the numbers
	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i] < listOne[j]
	})

	sort.Slice(listTwo, func(i, j int) bool {
		return listTwo[i] < listTwo[j]
	})

	var result []int
	for i, v := range listOne {
		if v < listTwo[i] {
			result = append(result, listTwo[i]-v)
		} else {
			result = append(result, v-listTwo[i])
		}
	}

	sum := getSum(result)
	fmt.Printf("Total distance between numbers: %d\n", sum)

	// PART TWO LOGIC
	firstArrCount := make(map[int]int)
	for _, valOne := range listOne {
		for _, valTwo := range listTwo {
			if valOne == valTwo {
				firstArrCount[valOne] = firstArrCount[valOne] + 1
			}
		}
	}

	var finalResult int
	for k, v := range firstArrCount {
		finalResult += k * v
	}

	fmt.Printf("Sum of count of repeated numbers %d", finalResult)

}
