package main

import "fmt"

func main(){
	var arr [3]int = [3]int{1,2,3,}
	fmt.Println(arr[1])

	arr2 := [3]int{4,5,6} // short decleration

	fmt.Println(arr2[1])

	// slices

	s := []int{8,9,10}

	s = append(s,90)

	fmt.Println("Slices Elelment" , s[3] , s)

	s1 := s[0:3]

	fmt.Println(s1[0])
}