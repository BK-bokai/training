package main // 可執行的程式必須使用main
import "fmt" // 載入內建的fmt 封包，用來做基本的輸入輸出
func main() {// 建立 main 函式，程式的進入點
	// 輸出訊息到終端: fmt.Println(資料,資料,...)

	/*
	資料型態
	*/
	fmt.Println(3) // int
	fmt.Println(3.14159) // float64
	fmt.Println("測試中文") // string
	fmt.Println(true) // bool
	fmt.Println('a') // 字符rune

	/*
	變數的使用
	*/
	// int
	var diameter int // 直徑
	diameter = 5 // 指定資料
	fmt.Println(diameter)

	// float64
	var pi float64=3.14159 // pi
	fmt.Println(pi)

	// 前轉型態
	fmt.Println(float64(diameter) * pi)

	// string
	var chat string = "hello"
	fmt.Println(chat)

	// bool
	var isBool bool = true
	fmt.Println(isBool)

	// rune
	var a rune = 'a'
	fmt.Println(a)
}