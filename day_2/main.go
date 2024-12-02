package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func readFile() [][]int {
	var result [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Cannot read file")
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		// read each line and split it
		line := sc.Text()

		collection := strings.Split(line, " ")

		integers := make([]int, 0, len(collection))

		for _, raw := range collection {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			integers = append(integers, v)
		}
		result = append(result, integers)

	}

	if err := sc.Err(); err != nil {
		log.Fatal("error scanning the file")
	}
	return result
}

func main() {
	nextIdx := 1
	var diff int
	var isIncreasing bool
	var isDecreasing bool
	var safeLevels int

	collection := readFile()

	for _, col := range collection {
		maxIdx := len(col) - 1

		for i, v := range col {
			fmt.Println(reflect.TypeOf(v))
			if i == maxIdx {
				nextIdx = 0
			}

			if i == maxIdx {
				fmt.Println("breaking because equeal index")
				break
			}

			if v == col[i+nextIdx] {
				diff = 0
				fmt.Printf("breaking because equeal values: %v-%v-%v\n", v, col[i+nextIdx], col)
				break
			}

			if v > col[i+nextIdx] {
				if isIncreasing {
					diff = 0
					fmt.Println("breaking because it started increasing")
					break
				}
				isDecreasing = true
				diff = v - col[i+nextIdx]
			}

			if v < col[i+nextIdx] {
				if isDecreasing {
					diff = 0
					fmt.Println("breaking because it started decreasing")
					break
				}
				isIncreasing = true
				diff = col[i+nextIdx] - v
			}

			if diff > 3 {
				diff = 0
				fmt.Println("breaking because the diff is too high")
				break
			}
		}

		isIncreasing = false
		isDecreasing = false
		nextIdx = 1
		if diff > 0 && diff <= 3 {
			safeLevels += 1
		}

	}

	fmt.Println(safeLevels)

}
