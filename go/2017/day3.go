package twentyseventeen

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Day3Solve(inputFile string) {
	file, err := os.Open(inputFile)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		p1 := day3(input)
		// p2 := Day3Bonus(input)
		fmt.Printf("d3p1 = %d\n", p1)
	}

}

// day3 Algorithm relies on square numbers. Designed for efficiency this algorithm avoids costly operations such as sqrt
// First, determine the floor of the root of the number-1. i.e. 12 -> 3, 21 -> 4, 1 -> 0
// Second, determine the numbers between i**2 and (i+1)**2 with the smallest manhattan distance,
// i.e. 1, 2, 11, 28 horizontally and 1, 4, 15, 34 vertically
// Third, find the minimum between the number and the v/h numbers and add that to the distance to the h/v number
func day3(goal float64) int {
	i := 0.0
	for ; math.Pow(i+1, 2) < goal; i++ {
	}

	v := math.Pow(i, 2) + math.Ceil(i/2)
	h := math.Pow(i+1, 2) - math.Floor(i/2)

	return int(math.Min(math.Abs(goal-v), math.Abs(goal-h)) + math.Ceil(i/2))
}

func day3Bonus(goal int) int {
	panic("not implemented yet")
}
