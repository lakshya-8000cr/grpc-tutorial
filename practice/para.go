package main

import "fmt"

func calculateArea(len , width float64) float64{
	area := len * width

	return area 
}


func divide(a , b float64) (float64 ,error) {
   if(b==0) {
	return 0 , fmt.Errorf("cannot be divided by 0")
   }

   return a /b , nil
}


func main() {
	fmt.Println(calculateArea(10 , 20))	
	fmt.Println(divide(10 , 20))
}