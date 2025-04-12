package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Number of ranks: ")
	scanner.Scan()
	numberOfRecord := scanner.Text()
	rankCount, _ := strconv.Atoi(numberOfRecord)

	fmt.Print("Current recorded ranks: ")
	scanner.Scan()
	currentRankRecorded := strings.Fields(scanner.Text())

	if len(currentRankRecorded) != rankCount {
		fmt.Println("Input rank less or more than " + strconv.Itoa(rankCount))
		return
	}

	scores := make([]int, rankCount)
	for i := 0; i < rankCount; i++ {
		scores[i], _ = strconv.Atoi(currentRankRecorded[i])
	}

	fmt.Print("Number of trials: ")
	scanner.Scan()
	numberOfTrials := scanner.Text()
	trialCount, _ := strconv.Atoi(numberOfTrials)

	fmt.Print("Point scored: ")
	scanner.Scan()
	trialPoint := strings.Fields(scanner.Text())

	if len(trialPoint) != trialCount {
		fmt.Println("Input score less or more than " + strconv.Itoa(trialCount))
		return
	}
}
