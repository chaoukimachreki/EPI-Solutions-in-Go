package main

import "fmt"
 const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Cyan   = "\033[36m"
    )


func multiplyBits(x, y uint) uint {
    var sum uint
   
    for x != 0 {
        if x&1 != 0 {
	    	
            sum = Add(sum, y)
	           }
        x >>= 1
        y <<= 1
    }
    return sum
}

func Add(a, b uint) uint {
    perm_a, perm_b := a, b
    var carryin, k, sum uint = 0, 1, 0

    for perm_a != 0 || perm_b != 0 {
        ak := a & k
        bk := b & k
        carryout := (ak & bk) | (ak & carryin) | (bk & carryin)
        sum |= ak ^ bk ^ carryin
        carryin = carryout << 1
        k <<= 1
        perm_a >>= 1
        perm_b >>= 1
    }

    return sum | carryin
}

func main(){
 
 	x:=uint(25)
	y:=uint(7)
	multiply:=multiplyBits(x,y)
	 fmt.Printf("%sThe result of the multiplication is: %s%d%s\n", Yellow, Green, multiply, Reset)
    fmt.Printf("%sAnd in binary: %s%064b%s\n", Yellow, Cyan, multiply, Reset)
}
