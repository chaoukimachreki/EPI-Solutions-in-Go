package main

import (
    "fmt"
)

func bitwiseDiv(x, y uint32) (uint32,error) {
    if y == 0 {return 0 , fmt.Errorf("Division by zero mate !") }
    var result uint32 = 0
    var power int = 32
    yPower := uint64(y) << power

    
    for x >= y {
	for yPower > uint64(x) {

		yPower >>=1
		power --
	} 
	if power >=32 {
		return 0 ,fmt.Errorf("power overflow")
	}
	result += 1<<power
	x-=uint32(yPower)

	
    }
    return result, nil
}

func main() {
    x := uint32(30)
    y := uint32(0)

    quotient,err := bitwiseDiv(x, y)
    if err != nil {
	fmt.Println("hold up ! you messed something up look : ",err)
    }

    fmt.Printf("Quotient of %d / %d = %d\n", x, y, quotient)
}

