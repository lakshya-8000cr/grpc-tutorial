package main

import "fmt"

type Service struct {
	Port int 
	Status string
}

func main(){
    m := map[string]int{
		"one" : 1 ,
		"two" : 2,
		"three" : 3,
		"four" : 4,
	}

	service := map[string]Service{
	"api" : {Port : 83 , Status : "Running"},
	"worker" : {Port :90 , Status : "Stopped"},
}


	for key , val := range m {
		fmt.Println(key,val)
	}

	val , ok := m["one"]

	if ok {
      fmt.Println("found" , val)
	} else{
	fmt.Println("not found")	
	}


	for key , val := range service{
		fmt.Println(key , val.Status)
	}


}