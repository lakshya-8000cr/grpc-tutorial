package main 

import "fmt"

func filetEven(arr []int) []int {
	var  slice []int

	for _, vl := range arr {
		if vl % 2 == 0 {
			slice = append(slice , vl)
		}
	}

	return slice 

}


func main() {
	result := []int{1,2,3,4,5,6,7,8}

	res := filetEven(result)

	fmt.Println(res)
}