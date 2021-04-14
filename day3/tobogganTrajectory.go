// tobogganTrajectory uses gofuncs to solve the 2020 day 3 problem. It'd be simpler to do this in one thread, but this was a neat way to learn about go concurrency and the issues one runs into
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// replace rune i of string in with rune r
// for debugging
func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	// store this character as a byte for easy comparison
	treeChar := byte('#')
	// store the integer changes in x and y per unit time
	dx := 3
	dy := 1

	// tree counting helper constructs
	treesHit := 0
	treeChan := make(chan interface{})     // to signal that a tree was hit
	stopCounting := make(chan interface{}) // to signal that we can stop counting
	countTrees := func() {
		for {
			select {
			case <-treeChan:
				treesHit++
			case <-stopCounting:
				stopCounting <- 1
				return
			}
		}
	}

	// a gofunc will be spun up for each channel that waits for a start position, sees if that is a tree on that slopeLine, then passes the next position off to the next gofunc
	var slopeLines []string
	var channels []chan int
	goDownSlopeAtLine := func(slopeAndChannelIndex int) {
		for {
			startPos := <-channels[slopeAndChannelIndex]
			startPos = startPos % len(slopeLines[slopeAndChannelIndex])
			if slopeLines[slopeAndChannelIndex][startPos] == treeChar {
				// slopeLines[slopeAndChannelIndex] = replaceAtIndex(slopeLines[slopeAndChannelIndex], 'X', startPos)
				treeChan <- 1
			}
			// fmt.Println(slopeLines[slopeAndChannelIndex])
			channels[slopeAndChannelIndex+dy] <- startPos + dx
		}
	}

	// read in data and spin up channels and slopeLines
	scanner := bufio.NewScanner(f)
	slopeNumber := 0
	for scanner.Scan() {
		thisLine := string(scanner.Text())
		slopeLines = append(slopeLines, strings.TrimSpace(thisLine))
		channels = append(channels, make(chan int))
		go goDownSlopeAtLine(slopeNumber)
		slopeNumber++
	}

	// create the final channel we'll pull the end-signal from
	finalChannel := make(chan int)
	channels = append(channels, finalChannel)

	go countTrees()

	channels[0] <- 0

	<-finalChannel
	// send signal and MAKE SURE it was done
	stopCounting <- 1
	<-stopCounting

	fmt.Println(treesHit)
}
