package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Flag for debugging
const debug = true

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

	// Variables we're interested in
	maxCalories := 0
	maxCaloriesElfId := 0

	currentElfId := 1
	currentElfCalories := 0

	// Won't work if file has only one line
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

		if currentElfCalories > maxCalories {
			log(fmt.Sprintf("Updating max calories, new max: %d (elf id: %d)", currentElfCalories, currentElfId))

			maxCalories = currentElfCalories
			maxCaloriesElfId = currentElfId
		}

		currentElfId++
		currentElfCalories = 0
	}

	fmt.Println(fmt.Sprintf("Elf with most calories is %d with %d calories", maxCaloriesElfId, maxCalories))
}
