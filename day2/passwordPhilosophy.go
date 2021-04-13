package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type password struct {
	character string
	min       int
	max       int
	password  string
}

func (p *password) verify() bool {
	count := strings.Count(p.password, p.character)
	return count <= p.max && count >= p.min
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	var data []password
	count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parsedLine := regexp.MustCompile(`(\d+)-(\d+) (\w): (.*)`).FindSubmatch([]byte([]byte(scanner.Text())))

		character := string(parsedLine[3])
		min, err := strconv.Atoi(string(parsedLine[1]))
		check(err)
		max, err := strconv.Atoi(string(parsedLine[2]))
		check(err)
		passphrase := string(parsedLine[4])

		p := password{
			character: character,
			min:       min,
			max:       max,
			password:  passphrase,
		}

		data = append(data, p)
	}

	for _, p := range data {
		if p.verify() {
			count++
		}
	}

	fmt.Println(count)
}
