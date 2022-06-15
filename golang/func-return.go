package main
import "fmt"

func main() {
	var resultStr string
	var result int
	resultStr, result = add(2, 3)
	fmt.Println(resultStr, result)
}

func add(n1 int, n2 int) (string, int) {
	var result int=n1+n2
	return  fmt.Sprintln(n1, "+", n2, "="), result
}
