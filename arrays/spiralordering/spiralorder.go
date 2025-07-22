package main 

import ( "fmt"
	"errors"
	)

func spiralOrdering(matrix [][]int)([]int,error){
	numRows := len(matrix)
	if numRows == 0  {return nil,errors.New("matrix cant be empty")}
	numColumns := len(matrix[0])
	totalSize := numRows * numColumns
	result := make([]int,0,totalSize)
	direction := [4][2]int{{0,1},{1,0},{0,-1},{-1,0}}
	x,y,dir := 0,0,0
	for i :=0; i < totalSize; i++ {

		result = append(result,matrix[x][y])
		matrix[x][y] = 0
		next_x := x + direction[dir][0]
		next_y := y + direction[dir][1]

		if next_x < 0 || next_x >= numRows || next_y < 0 || next_y >= numColumns || matrix[next_x][next_y] == 0 {
 
				
				dir = (dir + 1) % 4
				next_x = x + direction[dir][0]
				next_y = y + direction[dir][1]

		}
		x, y = next_x, next_y

	}
	return result,nil


}

func main(){
	
	matrix := [][]int{
       // {1, 2, 3},
       // {4, 5, 6},
       // {7, 8, 9},
	//{19, 28, 39},

    }
    	result, err := spiralOrdering(matrix)
	if err != nil {

		fmt.Println("error found : ",err)
		return
	}

	fmt.Println("final result : ",result)



}
