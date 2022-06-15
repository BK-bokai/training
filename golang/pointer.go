package main
import "fmt"

func main () {
	// 建立存放資料的變數
	// 取得記憶體位址: &變數名稱
	// 存放到指標變數: *資料型態
	// 反解指標變數: *指標變數

	var x int = 5
	fmt.Println("原來的資料", x)

	// 印出x的pointer
	fmt.Println("資料的記憶體位址", &x)

	// 儲存x的pointer
	var xPtr *int = &x
	// 利用指標取的內容
	fmt.Println("反解指標變數" ,*xPtr)

}