package main
import "fmt"

func main() {
	var x int = 10
	// var xPtr *int = &x
	// pass by Pointer
	add(&x, 20)

	fmt.Println(x)
}

func add(xPtr *int, add_num int) {
	*xPtr += add_num;
}