package main
import "fmt"
func main() {
	var x float64 = 3
	var y float64 = 4

	// 算術運算+ - * /
	fmt.Println(x+y, x-y, x*y, float64(x/y))

	// 指定運算 += -+ *= /=
	x += 1
	fmt.Println(x) // 4
	x -= 1
	fmt.Println(x) // 3
	x *= 5
	fmt.Println(x) // 15
	x /= 3
	fmt.Println(x) // 5

	// 單元運算 ++ --
	x++
	fmt.Println(x) // 6
	x--
	fmt.Println(x) // 5

	// 比較運算 >, <, >=, <=, ==
	fmt.Println(x > y)// true

	// 邏輯運算 !, &&, ||
	fmt.Println(!(x > y)) // false
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true



	


}