package main
import "fmt"

type Rectangle struct {
    X, Y          int
    Width, Height int
}

func (r Rectangle) IsValid() bool {
    return r.Width >= 0 && r.Height >= 0
}

func IsIntersecting(a, b Rectangle) bool {
    return a.X < b.X+b.Width &&
        a.X+a.Width > b.X &&
        a.Y < b.Y+b.Height &&
        a.Y+a.Height > b.Y
}

func Intersect(a, b Rectangle) (Rectangle, bool) {
    if !IsIntersecting(a, b) {
        return Rectangle{}, false
    }

    left := max(a.X, b.X)
    top := max(a.Y, b.Y)
    right := min(a.X+a.Width, b.X+b.Width)
    bottom := min(a.Y+a.Height, b.Y+b.Height)

    return Rectangle{
        X:      left,
        Y:      top,
        Width:  right - left,
        Height: bottom - top,
    }, true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    rect1 := Rectangle{X: 0, Y: 0, Width: 5, Height: 5}
    rect2 := Rectangle{X: 3, Y: 3, Width: 4, Height: 4}

    if intersection, ok := Intersect(rect1, rect2); ok {
        fmt.Printf("Intersection found: %+v\n", intersection)
    } else {
        fmt.Println("No intersection found.")
    }
}
