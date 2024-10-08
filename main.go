package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Rule int

const (
	kill Rule = iota
	alive
	keep
)

func initGrid(size int) [][]byte {
	grid := make([][]byte, size)

	for row := 0; row < size; row++ {
		grid[row] = make([]byte, size)
		for col := 0; col < size; col++ {
			grid[row][col] = byte(rand.Intn(2))
		}
	}

	return grid
}

func getNeighbours(grid [][]byte, size, col, row int) int {
	leftbound := -1
	rightbound := 1
	upperbound := -1
	lowerbound := 1

	// If we are on the first row, we do not check the upper row
	if row == 0 {
		upperbound = 0
	}

	// If we are on the last row, we do not check the lower row
	if row == size-1 {
		lowerbound = 0
	}

	// If we are on the first column we do not check the left column
	if col == 0 {
		leftbound = 0
	}

	// If we are on the last column we do not check the right column
	if col == size-1 {
		rightbound = 0
	}

	count := 0
	for i := upperbound; i <= lowerbound; i++ {
		for j := leftbound; j <= rightbound; j++ {
			// If the offsets are both 0 it means we are on the cell, thus we skip
			if i == 0 && j == 0 {
				continue
			}

			count += int(grid[row+i][col+j])
		}
	}

	return count
}

// Moore neighbourhood
func checkRules(grid [][]byte, size, col, row int) Rule {
	n := getNeighbours(grid, size, col, row)

	// The neighbour number
	if n >= 4 || n <= 1 {
		return kill
	} else if n == 3 {
		return alive
	} else {
		return keep
	}
}

func updateGrid(grid [][]byte) {

	size := len(grid)
	// We create the second matrix in which we have the updated values
	updated := make([][]byte, size)
	for row := 0; row < size; row++ {

		// We create each row
		updated[row] = make([]byte, size)

		for col := 0; col < size; col++ {
			r := checkRules(grid, size, col, row)

			switch r {
			case kill:
				updated[row][col] = 0
			case alive:
				updated[row][col] = 1
			default:
				updated[row][col] = grid[row][col]
			}
		}
	}

	// We update the original matrix with the updated values
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j] = updated[i][j]
		}
	}
	printGrid(grid)
}

// We print the grid
func printGrid(grid [][]byte) {
	size := len(grid)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			fmt.Printf("%d ", grid[row][col])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// The terminal codes to clear the screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	tbp := flag.Int("t", 1000, "The time in milliseconds taken between each print")
	gs := flag.Int("s", 5, "The size of the matrix")

	flag.Parse()

	if *tbp <= 0 {
		fmt.Printf("Cannot have less than 0 ms of time between prints\n")
		return
	}

	if *gs <= 3 {
		fmt.Printf("The size of the matrix cannot be smaller than 4\n")
		return
	}

	clearScreen()
	grid := initGrid(*gs)

	for {
		clearScreen()
		updateGrid(grid)
		time.Sleep(time.Duration(*tbp) * time.Second)
	}

}
