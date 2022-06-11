package main
import "fmt"
func main() {
	var x int
	var y int

	fmt.Println("請輸入兩個數字用空格隔開")
	fmt.Scanln(&x, &y)
	if x > y {
		fmt.Printf("%d > %d", x, y)
	} else if x < y {
		fmt.Printf("%d < %d", x, y)
	} else {
		fmt.Printf("%d = %d \n", x, y)
	}
}