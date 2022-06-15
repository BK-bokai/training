package main
import "fmt"
func main() {
	var intArr [3]int
	var strArr [3]string=[3]string{"php", "golang", "java"}
	var boolArr [3]bool=[3]bool{false, true, true}

	intArr[0] = 14
	intArr[1] = 19
	intArr[2] = 12
	// intArr[3] = 1 //array index 3 out of bounds [0:3]
	fmt.Println(intArr, strArr, boolArr)

	var index int
	for index = 0; index <len(intArr); index++ {
		fmt.Println(index)
	}

	for index, element := range strArr {
		fmt.Println(index, element)
	}
}