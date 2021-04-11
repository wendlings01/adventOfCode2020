package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// naive/brute force, fine for the single use case
// but with the addition of wanted 2020 from 3 numbers we should do better
func find2020(array *[]int) int {
	for i := 0; i < len(*array)-1; i++ {
		restOfTheNumbers := (*array)[i:len(*array)]
		for j := 0; j < len(restOfTheNumbers); j++ {
			if (*array)[i]+restOfTheNumbers[j] == 2020 {
				return (*array)[i] * restOfTheNumbers[j]
			}
		}
	}

	// probably not great golang style, maybe this should return an error as well as an int?
	return -1
}

// some assumptions are made about the reliability of input, could be more robust
// does data need to be a pointer? maybe? I don't know how I feel about the data duplication, but it might be performant enough
// if this is reused, consider:
//		abstracting numbersSoFar out of the parameters and pass it to a helper function that recurses
//		simplify the base case(additionsNeeded == 1) by replacing numbersSoFar with currentSum and currentProduct parameters
//		maybe do the math before the loop and after "error" conditions. this may not play nice with the end of the slice though
func addToFind2020(additionsNeeded int, numbersSoFar []int, data *[]int) int {
	if additionsNeeded < 1 {
		return -1
	}

	for i := 0; i < len(*data); i++ {
		if additionsNeeded == 1 {
			var sum int
			for _, val := range numbersSoFar {
				sum += val
			}
			if sum+(*data)[i] == 2020 {
				product := (*data)[i]
				for _, val := range numbersSoFar {
					product *= val
				}
				return product
			}
		} else {
			restOfTheNumbers := (*data)[i:len(*data)]

			result := addToFind2020(additionsNeeded-1, append(numbersSoFar, (*data)[i]), &restOfTheNumbers)
			if result != -1 {
				return result
			}
		}
	}

	return -1
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	var data []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		data = append(data, i)
	}

	// result := find2020(&data)

	result := addToFind2020(3, make([]int, 0), &data)

	fmt.Println(result)
}
