package main

import "fmt" 

func add(a int , b int) int {
	sum := a + b;

	fmt.Println(sum)

	return sum
}

func main() {
	result := add(40 , 30)
	fmt.Println("result is : " , result);
}

