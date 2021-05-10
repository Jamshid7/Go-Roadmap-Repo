package main
import "fmt"
func main() {
	fmt.Println(oddEven(4))
	fmt.Println(sum(7, 4))
}

func oddEven(a int) string{
	var res string
	if(a%2==0){
		res = "even"
	}else{
		res = "odd"
	}
	return res
}

func sum(a, b int) int{
	sum := a+b
	return sum
}
