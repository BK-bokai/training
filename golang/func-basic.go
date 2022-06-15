package main
import "fmt"
func main() {
	say("Hello")
	say("I'm Aaron")
	add(2,3)
	sum(10)

}

func say(msg string) {
	fmt.Println(msg)
}

func add(n1 int, n2 int) {
	var result int= n1+n2
	fmt.Println(result)
}

func sum(max int) {
	num := 0
	var result int
	for num =1; num<=max; num+=1 {
		result += num
	}
	fmt.Println(result)
}