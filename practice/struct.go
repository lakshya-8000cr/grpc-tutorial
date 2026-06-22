package main
import "fmt"

type User struct {
	Name string
	age int
}



func main() {
	u1:= User{Name:"Lakshya Chauhan", age:29}
u2:=User{Name:"Manu" , age:20}
if u1.age > u2.age {
	fmt.Println(u1.Name , " age is greter than " , u2.Name )
}else{
	fmt.Println("f")
}
}