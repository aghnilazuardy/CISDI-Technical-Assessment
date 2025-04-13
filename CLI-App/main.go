package main

import (
	"bufio"
	"cisdi-technical-assessment/helper"
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

	records := make([]int, rankCount)
	for i := 0; i < rankCount; i++ {
		records[i], _ = strconv.Atoi(currentRankRecorded[i])
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

	scores := make([]int, trialCount)
	for i := 0; i < trialCount; i++ {
		scores[i], _ = strconv.Atoi(trialPoint[i])
	}

	// calculting rank of scored point
	result := helper.CalculateRank(records, scores)

	// print result
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
