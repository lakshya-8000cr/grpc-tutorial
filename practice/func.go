package main

import "fmt" 

func add(a int , b int) int {
	sum := a + b;

	fmt.Println(sum)

	return sum
}

func divide(a , b float64) (float64 ,error) {
   if(b==0) {
	return 0 , fmt.Errorf("cannot be divided by 0")
   }

   return a /b , nil
}
func main() {
	result := add(40 , 30)
	fmt.Println("result is : " , result);

	fmt.Println(divide(10 ,20))
}

