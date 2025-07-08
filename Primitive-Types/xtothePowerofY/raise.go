package main

import (
	"fmt"
	"errors"
)

func Power(x float64,y int64)(float64,error){
	//edge cases , when x is zero and (y is zero or negative)
	
	if x == 0 {
		if y == 0 {
			return 0, errors.New("hold up rasing 0 to zero ? drunk ?")
			}
		if y<0 {
			return 0, errors.New("what the heck ? raising zero to a negative value ? are you nuts?")
		}
		return 0, nil

	}

	negativeExponent := y < 0
	exponent := y
	result := 1.0

	if negativeExponent {
		exponent = -exponent 
		x = 1.0/x
	}

	for exponent != 0 {

		if exponent & 1 != 0 {

			result *= x
		}

		exponent >>=1
		x*=x
	}
	return result, nil


}

func main() {

	cases := []struct {
		x float64
		y int64
	}{
	
	{2.0, 10},
        {2.0, -3},
        {5.0, 0},
        {0.0, 5},
        {0.0, -2},
        {1.0001, 1000000},

	}

	for _, kase := range cases {
		
		base := kase.x
		power := kase.y
		result,err := Power(base,power)
		if err != nil {
			fmt.Printf("Pow(%g, %d): error → %v\n", kase.x, kase.y, err)
		
		} else {
			fmt.Printf("Pow(%g, %d) = %.10g\n", kase.x, kase.y, result)
		}
	}
	


}
