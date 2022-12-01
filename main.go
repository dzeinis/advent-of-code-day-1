package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// Flag for debugging
const debug = false

// How many elves from the top should we include
const numberOfElvesToCount = 3

// Helper for debugging
func log(message string) {
	if !debug {
		return
	}

	fmt.Println(message)
}

func main() {
	// Get the input file
	fileName := os.Args[1]

	log(fmt.Sprintf("Reading from file [%s]", fileName))

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	// Close the file at the end of the function
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfCalories []int

	currentElfCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		log(fmt.Sprintf("Read line [%s]", line))

		if line != "" && line != "\n" {
			parsedCalories, err := strconv.Atoi(line)

			if err != nil {
				panic(err)
			}

			currentElfCalories += parsedCalories

			continue
		}

		elfCalories = append(elfCalories, currentElfCalories)
		currentElfCalories = 0
	}

	// Add the last elf, as the loop will exit before processing the last one
	// Also check if it is non-zero
	if currentElfCalories > 0 {
		elfCalories = append(elfCalories, currentElfCalories)
	}

	// Sort the elves and their calories
	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	log(fmt.Sprintf("Sorted list: %v", elfCalories))

	if len(elfCalories) == 0 {
		fmt.Printf("No elves provided")

		return
	}

	caloriesForTopElves := 0

	numberOfElves := int(math.Min(numberOfElvesToCount, float64(len(elfCalories))))

	for j := 0; j < numberOfElves; j++ {
		caloriesForTopElves += elfCalories[j]
	}

	fmt.Printf("Top %d elves have %d calories", numberOfElves, caloriesForTopElves)
}
