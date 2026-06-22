package main

import "fmt"


func main(){

	arr := []string{"api" , "worker", "scheduler"}

	for i , name := range arr {
          fmt.Println(i , name);
	}

}