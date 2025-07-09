package main

import (
    
    "fmt"
)

// Color defines the available colors used in the Dutch Flag problem.
type Color int

const (
    Red Color = iota
    Green
    Blue
)

// String provides the string representation for a Color.
func (c Color) String() string {
    switch c {
    case Red:
        return "red"
    case Green:
        return "green"
    case Blue:
        return "blue"
    default:
        return "unknown color"
    }
}

// DutchFlag reorders the input slice so that all elements less than, equal to,
// and greater than the pivot are grouped together respectively.
//
// It returns an updated slice or an error if the pivot index is out of bounds.
func DutchFlag(pivotIndex int, colors []Color) ([]Color, error) {
    if pivotIndex < 0 || pivotIndex >= len(colors) {
        return nil, fmt.Errorf("invalid pivot index: %d", pivotIndex)
    }

    pivot := colors[pivotIndex]
    smaller, equal, larger := 0, 0, len(colors)

    for equal < larger {
        switch {
        case colors[equal] < pivot:
            colors[smaller], colors[equal] = colors[equal], colors[smaller]
            smaller++
            equal++
        case colors[equal] == pivot:
            equal++
        default: // colors[equal] > pivot
            larger--
            colors[equal], colors[larger] = colors[larger], colors[equal]
        }
    }
    return colors, nil
}

func main() {
    colors := []Color{Red, Blue, Green, Blue, Blue, Green}
    pivotIndex := 2

    sortedColors, err := DutchFlag(pivotIndex, colors)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    for _, c := range sortedColors {
        fmt.Printf("%s ", c)
    }
    fmt.Println()
}

