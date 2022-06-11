package main
import "fmt"

func main() {
	var x int
	fmt.Println("請輸入一個數字:")
	fmt.Scanln(&x)

	for x > 0 {
		// fmt.Println(x)
		x--
	}

	var result int
	for x = 1;x <= 50 ; x++ {
		result += x
	}
	// fmt.Println(result)

	var y int = 0
	var z int = 0
	for true {
		fmt.Println("請輸入int 由我幫你加總，若輸入0 將停止計算")
		fmt.Scanln(&z)
		
		if z == 0 {
			break
		}
		y += z
		fmt.Println(y)
	}

	
}