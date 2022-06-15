package main
import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	var Aaron Person=Person{"Aaron", 18}
	var Perla Person=Person{age:30, name:"Perla"}
	fmt.Println(Aaron.name, Aaron.age)
	Perla.name = "PaPa"
	fmt.Println(Perla.name, Perla.age)
}