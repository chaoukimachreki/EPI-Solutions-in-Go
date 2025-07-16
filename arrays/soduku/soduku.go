package main

import (
    "fmt"
    "math"
)

// IsValidSudoku checks if a partially filled Sudoku board has any conflicts.
// Zero represents an empty cell.
func IsValidSudoku(board [][]int) bool {
    size := len(board)

    // Check rows
    for i := 0; i < size; i++ {
        if HasDuplicate(board, i, 0, i+1, size) {
            return false
        }
    }

    // Check columns
    for j := 0; j < size; j++ {
        if HasDuplicate(board, 0, j, size, j+1) {
            return false
        }
    }

    // Check subgrids
    regionSize := int(math.Sqrt(float64(size)))
    for I := 0; I < regionSize; I++ {
        for J := 0; J < regionSize; J++ {
            if HasDuplicate(board,
                I*regionSize,
                J*regionSize,
                (I+1)*regionSize,
                (J+1)*regionSize) {
                return false
            }
        }
    }

    return true
}

// HasDuplicate returns true if any non-zero value appears more than once
// in the subgrid defined by [startRow:endRow) x [startCol:endCol)
func HasDuplicate(board [][]int, startRow, startCol, endRow, endCol int) bool {
    size := len(board)
    isPresent := make([]bool, size+1)

    for i := startRow; i < endRow; i++ {
        for j := startCol; j < endCol; j++ {
            val := board[i][j]
            if val != 0 {
                if isPresent[val] {
                    return true
                }
                isPresent[val] = true
            }
        }
    }

    return false
}

func main() {
    board := [][]int{
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9},
    }

    valid := IsValidSudoku(board)
    fmt.Println("Is the Sudoku valid?", valid)
}

