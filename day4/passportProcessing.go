package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) isValid() bool {
	if len(p.byr) == 0 {
		return false
	}
	if len(p.iyr) == 0 {
		return false
	}
	if len(p.eyr) == 0 {
		return false
	}
	if len(p.hgt) == 0 {
		return false
	}
	if len(p.hcl) == 0 {
		return false
	}
	if len(p.ecl) == 0 {
		return false
	}
	if len(p.pid) == 0 {
		return false
	}
	return true
}

func (p *passport) hasInfo() bool {
	if len(p.byr) > 0 {
		return true
	}
	if len(p.iyr) > 0 {
		return true
	}
	if len(p.eyr) > 0 {
		return true
	}
	if len(p.hgt) > 0 {
		return true
	}
	if len(p.hcl) > 0 {
		return true
	}
	if len(p.ecl) > 0 {
		return true
	}
	if len(p.pid) > 0 {
		return true
	}
	if len(p.cid) > 0 {
		return true
	}
	return false
}

func (p *passport) enterData(key, data string) {
	switch key {
	case "byr":
		p.byr = data
	case "iyr":
		p.iyr = data
	case "eyr":
		p.eyr = data
	case "hgt":
		p.hgt = data
	case "hcl":
		p.hcl = data
	case "ecl":
		p.ecl = data
	case "pid":
		p.pid = data
	case "cid":
		p.cid = data
	default:
		log.Fatal(key, data)
	}
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var passports []passport
	currentPassport := passport{}
	for scanner.Scan() {
		parsedLine := regexp.MustCompile(`(\w{3}):([^ ]*)`).FindAllSubmatch([]byte(scanner.Text()), -1)
		// fmt.Println(scanner.Text())
		// fmt.Printf("%q\n", parsedLine)
		if len(parsedLine) < 1 {
			// fmt.Printf("END OF PP\n")
			if currentPassport.hasInfo() {
				passports = append(passports, currentPassport)
				// fmt.Printf("Got pp: %q\n", currentPassport)
				// fmt.Println("#####################################")
			}
			currentPassport = passport{}
		} else {
			for _, data := range parsedLine {
				currentPassport.enterData(string(data[1]), string(data[2]))
			}
		}
	}
	if currentPassport.hasInfo() {
		passports = append(passports, currentPassport)
	}

	validPassports := 0
	for _, p := range passports {
		// fmt.Printf("%q\n", p)
		if p.isValid() {
			validPassports++
		} else {
			// fmt.Printf("INVALID: %q\n", p)
		}
	}

	fmt.Println(len(passports))
	fmt.Println(validPassports)
}
