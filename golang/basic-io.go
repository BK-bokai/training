package main
import "fmt"
func main() {
	var x int = 3
	var y int = 6

	// Println 結尾自帶換行符
	fmt.Println(x, y)
	// Print 會與下一個print在同一行
	fmt.Print(x, y)
	fmt.Print(x, y, "\n")
	

	// &變數名稱，取得變數指標(Pointer)
	fmt.Print("請輸入兩個數字用空格隔開:\n")
	fmt.Scanln(&x, &y)

	fmt.Println(x, "X" ,y, "=", x*y)
	
	/*
	%d: digit   (10進位的數字)
	%c: char    (字元)
	%s: string  (字串)
	%v: value   (值)
	%+v 見下方
	%#v 見下方
	*/
	var str string = "計算結果為:"
	fmt.Printf("%s : %d X %d = %d \n", str, x, y, x*y)
}