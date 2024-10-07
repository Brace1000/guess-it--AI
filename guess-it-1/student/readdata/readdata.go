package readdata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess/guessrange"
)

// ReadData reads numbers from standard input and predicts the range for the next number.
func ReadData() {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		// Convert input to float64
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		// Add the number to the list
		numbers = append(numbers, num)

		// Calculate the range for the next number prediction
		if len(numbers) > 1 {
			// Trim the length of numbers to be at most 2 by removing the first element
			if len(numbers) == 3 {
				numbers = numbers[1:]
			}
			lowerLimit, upperLimit := guessrange.Guessrange(numbers)
			fmt.Printf("%d %d\n", int(lowerLimit), int(upperLimit))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed reading input: %v\n", err)
	}
}
